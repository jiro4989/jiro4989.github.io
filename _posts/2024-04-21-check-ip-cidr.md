---
layout: default
title: "大量の IP が特定の CIDR に含まれているか調べるツールを作った"
date: 2024-04-21 09:00:00 +0900
categories: tech
---

# 大量の IP が特定の CIDR に含まれているか調べるツールを作った

アクセスログとかに書かれている大量のIPアドレスが、特定のCIDRに含まれるかどうかを高速にチェックしたくなったので、専用CLI checkcidr を作った。
使用言語は Go 。

* リポジトリ: <https://github.com/jiro4989/checkcidr>

たぶん誰かすでに作ってると思うんだけれど、最近コード書いてなかったし、気晴らしも兼ねて自分でガリガリ書いた。

* Table of contents
{:toc}

## 背景

アクセスログにはアクセス元IPアドレスが記録されており、それである程度アクセス元が特定できる。
サービスを運用してるとほぼ毎日なんらかの不正アクセスが試行されている。

で、怪しいアクセスをまとめてブロックできたらいいんだけれど、誤ってブロックしてはいけないものもブロックするとまずい。
例えば Google のクローラー。

Google のクローラーは、Web サイトを定期的にクローリングしている Bot です。
このクローラーの目的は、サイトをなんらかの方法で評価して、検索結果上位に載せること。
このクローラーを誤ってブロックしてしまうと、下手したら検索エンジン上にサイトが出てこなくなって SEO が悪化する可能性がある。

Google のクローラーからのアクセスの場合は UserAgent もそれとわかる名前になっている。
なので UserAgent でブロックしないようにしても良いんだけれど、攻撃者が UserAgent を Googlebot に偽装する可能性も 0 ではない。

Googlebot のサクセス元 CIDR は公開されているので、この CIDR のアクセスを許可すれば誤ったブロックを防げる。

* [クローラーが Googlebot などの Google クローラーであることを確認する](https://developers.google.com/search/docs/crawling-indexing/verifying-googlebot?hl=ja)

このように、なんらかの CIDR に特定の IP アドレスが含まれるか調べたくなる機会は多々ある。
しかし、アクセスログに記録されるIPアドレスは膨大で、且つ検査したい CIDR も1つではなく、何十個も存在したりする。
必然的に CIDR 数 x IP アドレス数 の総当たりでチェックして、見つかったら break 、って感じにしないととんでもなく時間がかかる。

こういった処理をそれなりに効率よく処理できて、進捗も確認できて、いろんな書式（JSONとか）で出力できるツールが欲しかったので、今回作った。

## 使い方

CIDR だけ書いたファイルと、IP だけ書いたファイルを引数として渡す。

```cidr.list
1.2.0.0/16
1.3.0.0/16
```

```ip.list
1.2.1.2
1.2.1.3
1.3.1.2
1.3.1.3
```

実行すると、以下の結果が得られる。

```bash
$ ./checkcidr cidr.list ip.list
ip_file=ip.list cidr=1.2.0.0/16 ip=1.2.1.2 contains=true
ip_file=ip.list cidr=1.3.0.0/16 ip=1.2.1.2 contains=false
ip_file=ip.list cidr=1.2.0.0/16 ip=1.2.1.3 contains=true
ip_file=ip.list cidr=1.3.0.0/16 ip=1.2.1.3 contains=false
ip_file=ip.list cidr=1.2.0.0/16 ip=1.3.1.2 contains=false
ip_file=ip.list cidr=1.3.0.0/16 ip=1.3.1.2 contains=true
ip_file=ip.list cidr=1.2.0.0/16 ip=1.3.1.2 contains=false
ip_file=ip.list cidr=1.3.0.0/16 ip=1.3.1.2 contains=true
```

JSON 形式でも出力できる。

```bash
$ ./checkcidr -style json cidr.list ip.list | head -n 13

[
  {
    "ip_file": "ip.list",
    "cidr": "1.2.0.0/16",
    "ip": "1.2.1.2",
    "contains": true
  },
  {
    "ip_file": "ip.list",
    "cidr": "1.3.0.0/16",
    "ip": "1.2.1.2",
    "contains": false
  },
```

JSON配列の出力の場合、すべての行の処理が完了するまで出力できなくて不便なので、1行JSON（JSON Stream）として出力するオプションもつけた。
この形式は jq でも処理できるし、割とよくある方式だと思う。

```bash
$ ./checkcidr -style json_stream cidr.list ip.list
{"ip_file":"ip.list","cidr":"1.2.0.0/16","ip":"1.2.1.2","contains":true}
{"ip_file":"ip.list","cidr":"1.3.0.0/16","ip":"1.2.1.2","contains":false}
{"ip_file":"ip.list","cidr":"1.2.0.0/16","ip":"1.2.1.3","contains":true}
{"ip_file":"ip.list","cidr":"1.3.0.0/16","ip":"1.2.1.3","contains":false}
{"ip_file":"ip.list","cidr":"1.2.0.0/16","ip":"1.3.1.2","contains":false}
{"ip_file":"ip.list","cidr":"1.3.0.0/16","ip":"1.3.1.2","contains":true}
{"ip_file":"ip.list","cidr":"1.2.0.0/16","ip":"1.3.1.2","contains":false}
{"ip_file":"ip.list","cidr":"1.3.0.0/16","ip":"1.3.1.2","contains":true}
```

### 進捗表示

検査処理のカウントが2500回を超えるたびに `.` を出力する処理を入れている。
また、100,000回に到達すると改行する。

総当たりチェックをすると試行回数がとんでもない数になる場合があるので、1 回ずつ進捗を出すのではなく、数千回に1回だけ進捗を出すようにしている。
画面がドットだらけ改行だらけとかになるとノイズになってしまうので、それを避けるためです。

こういう進捗表示は、ノイズにならないように、でも動作していることが視覚的に分かる良いバランスのとこを見つける必要がある。
進捗表示のためだけに依存ライブラリを増やすのも嫌なので、基本的に簡単な進捗表示で良いときはいつもこのスタイルで実装している。

今回は一般公開の CLI で採用したが、社内専用ツールを実装するときにも使える手法なのでおすすめ。
これらは標準エラー出力へ書き込むようにしてるので、検査結果出力の邪魔にはならない。
また、オプションで進捗表示を消せるようにしている。

```bash
$ ./checkcidr testdata/cidr_1.txt testdata/ip_1.txt > /dev/null
........................................ 100000
........
```

## 実装時間

だいたい 3 時間ほど。

## まとめ

* CIDR に IP が含まれるか検査する CLI checkcidr を作った
* 大量の試行回数になることを見越して、進捗を出力するようにした
