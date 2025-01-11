---
layout: default
title: "Google Search Console にサイトマップが登録されないけれど気にせず放置している"
date: 2025-01-09 00:28:25 +0900
categories: tech
categoriesJP: 技術
---

# Google Search Console にサイトマップが登録されないけれど気にせず放置している

一応このブログも Google の検索に出てきてほしいのでサイトマップを Google Search Console に登録している。

しかし、登録してから 1 ヶ月以上放置しても、以下のように「取得できませんでした」となっている。

![送信されたサイトマップの現状](https://github.com/user-attachments/assets/f21f0596-cb9e-41db-aed6-ec7cebc3b741)

ググった感じ「時間経過で解消した」といった事例が散見されたので、放置していたのだが全然解消しなかった。
仕方なく重い腰を上げて調べた。

* Table of contents
{:toc}

## 結論

最初に結論をいうと、途中で調査をやめた。
そのため、サイトマップが登録されない理由は不明なまま。
ただし、理由もなく調査をやめたわけではなく、必要なさそうだったのでやめた。

[サイトマップについて - Google 検索セントラル](https://developers.google.com/search/docs/crawling-indexing/sitemaps/overview?hl=ja#do-i-need-a-sitemap)
によると、サイトマップ不要な場合があるらしい。

そもそもサイトマップはなぜ必要なのかというと、それも上の記事に説明があった。
ざっと説明を読んで要約すると、サイトマップが必要な理由は以下のとおり。

1. Google などの検索エンジンがサイトをクローリングする際に、重要なページを効率的に見つけるためのファイル
1. 大きなサイトで必要になる。
   新規ページが作られた際に、他のページからリンクが貼られていないと、クローラーがそのページにたどり着けなくなる。
   この問題は特に大きなサイトで起こり得る

逆に、サイトが小さくてページがすべて他ページからリンクされており、
必ずページへ到達できるようになっているなら、サイトマップは不要ということ。
特に、このブログは記事をすべてトップページに列挙する作りなので、孤立ページは生まれない。
そのため、先の記事のとおりサイトマップは不要となる。

## サイトマップが読み込まれてないけれど Google 検索にインデックスされている

加えて、いつの間にかこのブログは Google 検索でヒットするようになっていた。

![次郎の貝塚でググった結果](https://github.com/user-attachments/assets/f715be88-4fc4-4bfd-a9d7-d47750f2fd09)

サイトマップは未だに「取得できませんでした」状態のままだけれど、Google 検索でヒットする。
なので Google Search Console は今だに変なことになってるけれど、Google 検索にのせたいって目的は達成出来ている。
そのため、現状のままで特に問題がないから、調査を打ち切った。

まぁ、このブログは自分がスマホからサクッと見られて、転職活動時に他人も見られるならそれで十分だったんだが。

## 調べたこと

一応調べたこともメモしておく。

### サイトマップの仕様

サイトマップの構文的には問題ないはず。
[Google の XML サイトマップ](https://developers.google.com/search/docs/crawling-indexing/sitemaps/build-sitemap?hl=ja#xml)と
[Sitemaps XML format - sitemaps.org](https://www.sitemaps.org/protocol.html)のプロトコルも確認したが
特段問題なさそうだった。

sitemaps.org による必須要件は次のとおり。

1. Sitemap のすべてのデータ値は [entry-escape](https://www.sitemaps.org/protocol.html#escaping) されていること
1. ファイルは UTF-8 でエンコードされていること
1. `<urlset>` タグで開始して、`</urlset>` タグで終了すること
1. `<urlset>` タグ内にネームスペース（プロトコル標準）を指定すること
1. 親 XML タグとして、URL ごとに `<url>` 要素を含むこと
1. `<url>` 親タグごとに子要素として `<loc>` を含むこと

これ以外のタグはすべて任意。
つまり `<lastmod>` とかは別になくていい。

確認した限り、やはり sitemap.xml の仕様を満たしている。
まぁ [jekyll-sitemap](https://github.com/jekyll/jekyll-sitemap) プラグインを使っているだけなので
変なサイトマップになりようがないはずだが。

### ファイルタイプ

一応 content-type も確認したが、application/xml なので、これも意図したもの。
実は html ファイルとして認識されている、的なこともない。

```bash
$ curl -v https://jiro4989.github.io/sitemap.xml 2>&1 | grep content-type
< content-type: application/xml
```

### robots.txt

robots.txt からも sitemap.xml の URL は貼ってある。

```bash
$ curl https://jiro4989.github.io/robots.txt
Sitemap: https://jiro4989.github.io/sitemap.xml
```

### `<loc>` が絶対 URL になっているか

[サイトマップの作成と送信 - Google 検索セントラル](https://developers.google.com/search/docs/crawling-indexing/sitemaps/build-sitemap?hl=ja)では、
参照 URL のプロパティは完全修飾された絶対 URL でなければならないと書かれていた。
/mypage.html みたいな相対パスは使ってはいけないらしい。
これも確認したが、絶対 URL になっていたので問題なかった。

### ブログをドメイン直下ではなく /blog でホスト

たぶん関係ないと思っていたけれど、原因切り分けで blog ってリポジトリを作って GitHub Pages を公開した。
/blog/sitemap.xml を登録して Google Search Console に登録してみたが、同じく「取得できませんでした」だった。
なのでホストしてるパスは関係ない。

ざっと調べたことはこの程度。

まぁ知らん間に、ググってブログがヒットするようになってたので、とりあえず良し。

以上。
