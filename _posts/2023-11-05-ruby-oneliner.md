---
layout: default
title: "Rubyワンライナーの基本とユースケース"
date: 2023-11-05 20:47:25 +0900
categories: tech
---

# Rubyワンライナーの基本とユースケース

仕事でRubyワンライナーを書くことがあります。
主にちょっとしたテキスト加工や、調査をする時だけですが、頻度は多いです。

どうせなので、自分が良く使うワンライナーの使い方と、細かい挙動についてまとめます。

* Table of contents
{:toc}

## まえがき

### Rubyワンライナーを覚える意義

Rubyワンライナーを覚えておくと、武器が増えます。
特に、突発的に発生した緊急度の高い作業時に役立ちます。
一番役に立つのは、主に障害対応の時でしょうか。

Rubyワンライナーの構文はawkと非常に似てます。
awkもワンライナーが得意なスクリプト言語です。
Rubyは、そんなawkよりもたくさんの機能を標準で備えています。

awkで処理しづらいテキスト処理の時に、Rubyワンライナーは威力を発揮します。
僕がRubyワンライナーを使うケースは、主にCSV、JSON、日付を扱う時です。
いずれも、awkで扱う時は苦労します。

awkを使えると簡単な集計処理ならサクッと実装できるようになります。
同様にRubyを使えると、CSVやJSONの簡単な処理ならサクッと実装できるようになります。

使える道具は多い方が芸の幅が広くなるので、覚えとくと役に立つでしょう。
ただし、Rubyはサーバによってはデフォルトで入ってないことも多いので、
awkとRuby両方使えるようになったほうが良いです。

### 使い捨てる前提でワンライナーを書く

ワンライナーは、ようは使い捨てのプログラムです。

通常であれば丁寧にスクリプトとして実装するものを、ワンライナーとして書く。
障害対応が発生した時はまさにそんなタイミングですね。

* 今すぐ調査して修正しないといけない
* 一分一秒が惜しい

そんな瞬間にワンライナーを書きます。

多くの場合、ワンライナーでやりたい処理は「今の自分にだけ」必要な作業です。
だから、通常のプログラムと違って、雑に書くことが許容されます。
今書いて今すぐ実行する必要があるから、変数名すらもマジックナンバーじみた極小の名前をつけがちです。
保守性は非常に低く、翌日ワンライナーを見て処理が理解できないとかはザラですね。

でも、ワンライナーはそれでいいんです。
明日同じワンライナーが必要になったなら、明日また同じワンライナーを書けばいいだけなので。
ワンライナーは使い捨てのプログラムだから、明日のことは考えなくていい。
リポジトリ管理のことも考えなくていいし、リポジトリ管理してはいけないです。

リポジトリで管理するプログラムは、保守するプログラムです。
保守するプログラムは、ワンライナーで実装するべきではありません。
普通にスクリプトファイルとして丁寧に実装するべきです。
適切な変数名を設定し、可読性を意識して実装する必要があります。

ワンライナーはリポジトリ管理しないプログラムだから、使い捨てる前提で書きましょう。

## ワンライナー用のオプション

まえがきが長くなったけれど、ここからようやくRubyワンライナーの説明をします。
まずはRubyでワンライナーする際に役立つオプション周りから。

### コマンドライン引数でRubyスクリプトを実行できる `-e`

Rubyのスクリプトをコマンドライン引数で実行できます。
標準入力を扱わないのであればこれだけで良いです。

```bash
ruby -e 'puts 1'     # -> 1
ruby -e 'puts 5 * 2' # -> 10
```

### 標準入力を1行ずつ扱える `-n`

awkだと特にオプションを指定せずとも標準入力を扱えます。
`$0`という特殊変数に、標準入力が1行ずつセットされます。

```bash
⟩ seq 3 | awk '{print $0}'
1
2
3
```

Rubyで同様のことをするには `-n` をセットします。
`$_`という特殊変数に、標準入力が1行ずつセットされるようになります。

```bash
⟩ seq 3 | ruby -ne 'puts $_'
1
2
3
```

### 行末の改行文字を削除する `-l`

`-n` だけだと、実は標準入力を扱いづらいです。
なぜなら`$_`の末尾に改行文字もついてしまってるからです。

こんな普通に文字列結合してみるとよく分かります。
改行文字のあとに文字列が追加されている。

```bash
⟩ seq 3 | ruby -ne 'puts $_ + " sushi"'
1
 sushi
2
 sushi
3
 sushi
```

この改行文字を除外するには `-l` を指定します。

```bash
⟩ seq 3 | ruby -lne 'puts $_ + " sushi"'
1 sushi
2 sushi
3 sushi
```

### 標準入力を複数列の配列として扱う `-a` (と `-F`)

awkでは、標準入力で渡された文字列を`$0`変数に1行ずつセットします。
そして、空白で区切られた列データとして、`$1`, `$2`といった数値変数にもセットします。

```bash
⟩ seq 9 | paste -d ' ' - - -
1 2 3
4 5 6
7 8 9

⟩ seq 9 | paste -d ' ' - - - | awk '{print $1 " | " $2 " | " $3}'
1 | 2 | 3
4 | 5 | 6
7 | 8 | 9
```

これと同様の処理をするのにRubyでは`-a`を指定します。
特殊変数`$F`に、配列としてセットされるようになります。

```bash
⟩ seq 9 | paste -d ' ' - - - | ruby -lane 'puts $F[0]'
1
4
7

⟩ seq 9 | paste -d ' ' - - - | ruby -lane 'puts $F[1]'
2
5
8
```

列の区切り文字を指定するのは、awkと同じで`-F`で指定します。

```bash
⟩ seq 9 | paste -d ',' - - -
1,2,3
4,5,6
7,8,9

⟩ seq 9 | paste -d ',' - - - | ruby -F, -lane 'puts $F[1]'
2
5
8
```

注意点としては、区切り文字を`-F`に隣接させないといけません。
`-F,`はOKだが、`-F ,`はエラーになります。

```bash
⟩ seq 9 | paste -d ',' - - - | ruby -F , -lane 'puts $F[1]'
ruby: No such file or directory -- , (LoadError)
```

### 困ったらとりあえず `-lane` でいい

とりあえず`-lane`って指定しとけばいいです。
awkとほぼ同じ感覚で標準入力を扱えるのでおすすめ。

```bash
⟩ seq 9 | paste -d ' ' - - - | ruby -lane 'puts $F[1]'
2
5
8

⟩ seq 9 | paste -d ',' - - - | ruby -lane 'puts "<" + $_ + ">"'
<1,2,3>
<4,5,6>
<7,8,9>
```

## ユースケース

Rubyワンライナー用のオプションの説明をしたので、次はユースケースをいくつか説明する。
awkで良くやる集計から、Rubyの標準ライブラリを使った処理などです。

### 数値の合計

配列変数では map メソッドが使用できます。
map メソッドを使うことで、配列の要素1つずつを加工して、別の型に変更できます。
これを利用して数値型に変換すれば、数値の合計を算出できます。

```bash
⟩ seq -s ' ' 10
1 2 3 4 5 6 7 8 9 10

⟩ seq -s ' ' 10 | ruby -lane 'puts $F.map{|s| s.to_i}.sum'
55
```

前述の `s.to_i` では、String クラスの `to_i` メソッドを呼び出しています。
こういったメソッド呼び出しは `&:to_i` という具合に省略して書くこともできます。
ただし、引数が必要なメソッドだとこの書き方はできません。

```bash
⟩ seq -s ' ' 10 | ruby -lane 'puts $F.map(&:to_i).sum'
55
```

`seq -s ' ' 10` は、1行に半角スペース区切りで数値を並べます。
1行ずつ数値が渡されるケースに対応するなら、以下のように書きます。
awkと同様に、BEGIN、ENDブロックが使えます。
ただしセミコロンで区切る必要があります。

```bash
⟩ seq 10 | ruby -lane 'BEGIN{ a = [] }; a.append($_.to_i); END{ puts a.sum }'
55
```

awkであれば、いきなり未初期化変数に対して代入ができます。
Rubyだと、そういった雑なコードは書けません。
かならず変数の初期化が必要です。

```bash
⟩ seq 10 | awk '{ a += $0 } END{ print a }'
55

# 変数 a は初期化してないのでエラーになる
⟩ seq 10 | ruby -lane 'a.append($_.to_i); END{ puts a.sum }'
-e:1:in `<main>': undefined local variable or method `a' for main:Object (NameError)
```

なお「変数が未定義の時だけ代入する」といった式が書けます。
この書き方ならBEGINブロックを省略できます。

```bash
⟩ seq 10 | ruby -lane 'a ||= []; a.append($_.to_i); END{ puts a.sum }'
55
```

こういった記号を使ったググりにくい構文は
[Rubyで使われる記号の意味（正規表現の複雑な記号は除く） - Ruby 3.2 リファレンスマニュアル](https://docs.ruby-lang.org/ja/latest/doc/symref.html)
にまとまっています。

### 数値の平均

前述の数値の合計に対して、配列変数 `$F` の長さで割ってやれば平均が出せます。
この時、左辺か右辺のどちらかを浮動小数点数 Float 型に変換しておく必要があります。
整数型変数のメソッド `to_f` を使えば Float 型に変換できます。

```bash
⟩ seq -s ' ' 10 | ruby -lane 'puts $F.map{|s| s.to_i}.sum.to_f / $F.length'
5.5
```

整数を整数で割ると、戻り値が整数になってしまうからです。

```bash
⟩ seq -s ' ' 10 | ruby -lane 'puts $F.map{|s| s.to_i}.sum / $F.length'
5
```

### GroupByして合計

これもawkと違って明示的にHashMapの変数を宣言しないといけません。

```bash
cat << EOS | ruby -lane 'a ||= {}; a[$F[0]] ||= 0; a[$F[0]] += $F[1].to_i; END{ puts a }'
aaa 100
bbb 200
aaa 100
ccc 300
bbb 200
ccc 300
EOS

{"aaa"=>200, "bbb"=>400, "ccc"=>600}
```

### 特定の文字列が存在する行だけ出力

後置ifで正規表現を書けます。
`/正規表現/`と書くだけです。

```bash
⟩ seq 10 | ruby -lane 'puts $_ if ~ /[2468]/'

2
4
6
8

# awkだとこう書く
⟩ seq 10 | awk '/[2468]/'
2
4
6
8
```

ただし、こういった単純な正規表現での絞り込みのみであれば `grep` で済ませたほうが良いです。

```bash
⟩ seq 10 | grep -E '[2468]'
2
4
6
8
```

### CSVの2列目を取得

RubyではCSVクラスを使うことで、CSVを簡単に扱えます。

CSVはカンマ区切りのデータだけれど、セル内にもカンマが含まれ得ます。
カンマ混じりのCSVをawkだけで扱うのはかなり難しいですが、Rubyだと簡単に処理できます。

```bash
⟩ echo '"hello,world","sushi,maguro","konbanwa"' | ruby -rcsv -e 'CSV(STDIN).each{|a| puts a[1]}'
sushi,maguro
```

この時、Rubyのオプションは `-lane` ではなく `-e` だけ付与して、標準入力をまるごとCSVクラスに渡します。
CSVのセルには、改行も含まれ得るので1行ずつ処理するとCSVを正しく処理できないことがあります。

### JSONの特定の列を取得

RubyではJSONクラスを使うことで、JSONを簡単に扱えます。

最近JSON形式のログを扱うことが多いです。
監視系がJSONパーサーを標準で備えているケースが多くなり、
それに合わせてログもJSON形式にすることも多くなったからですね。

そういった監視系に合わせたJSON形式のログの場合、1行に1つのJSONを出力します。
つまり、ワンライナーでJSONを扱う場合も、1行ずつ処理するのでほぼ問題ありません。

例えば、1行1JSONのログから数値配列 `value` を取得して合算する例を示します。

```bash
cat << EOS | ruby -rjson -lane 'n ||= []; v = JSON.parse($_); n.concat(v["value"]); END{ puts n.sum }'
{"time":"2021-01-02T11:02:15Z","level":"INFO","value":[1,2,3]}
{"time":"2021-01-03T11:02:15Z","level":"INFO","value":[5,6]}
{"time":"2021-01-04T11:02:15Z","level":"INFO","value":[7,8,9]}
EOS
```

前述の例は「1行に1つJSONが存在する」ものでしたが「入力全体がJSON」の場合は、次のように実装します。
`JSON.parse` の引数は文字列型を受け付けるので、`STDIN` (IO) 型ではなく、`STDIN.read` を渡す必要があります。

```bash
cat << EOS | ruby -rjson -e 'v = JSON.parse(STDIN.read); puts v["value"]'
{
  "time":"2021-01-02T11:02:15Z",
  "level":"INFO",
  "value":[1,2,3]
}
EOS
```

### 特定の日付の期間のログを抽出

RubyではTimeクラスを使うことで、日付を簡単に扱えます。
日付文字列のパースや、日付オブジェクトを生成したりできます。

日付を扱えると、ログファイルで特定の日付以降のログだけ抽出といった処理が書けます。
もちろんタイムゾーンも指定できるため、時差を考慮した抽出も可能です。

```bash
⟩ cat << EOS | ruby -rtime -lane 'puts $_ if Time.new(2021, 2, 1, 9, 0, 0, "+09:00") < Time.parse($F[0])'
2021-01-01T11:12:30Z hello1
2021-01-02T11:12:30Z hello2
2021-01-05T11:12:30Z hello3
2021-02-01T00:00:00Z 0ji 0byou
2021-02-01T00:00:01Z 0ji 1byou
2021-02-05T20:23:30Z hello4
2021-02-11T01:23:30Z hello5
2021-02-19T12:33:30Z hello6
2021-03-02T06:23:30Z hello7
2021-03-13T13:56:30Z hello8
2021-03-21T09:41:30Z hello9
EOS

2021-02-01T00:00:01Z 0ji 1byou
2021-02-05T20:23:30Z hello4
2021-02-11T01:23:30Z hello5
2021-02-19T12:33:30Z hello6
2021-03-02T06:23:30Z hello7
2021-03-13T13:56:30Z hello8
2021-03-21T09:41:30Z hello9
```

## 標準ライブラリ

僕がRubyワンライナーで処理したくなるデータは、主にCSV、JSONそして日付の3つです。
これは標準ライブラリとして専用クラスが用意されているため、非常に扱いやすいです。
それ以外にも、Rubyは多数の標準ライブラリを備えています。

[ライブラリ一覧](https://docs.ruby-lang.org/ja/latest/library/index.html)に、標準で使えるライブラリが載っています。
ここのライブラリの名前だけでも覚えておけば、有事の際に役立つでしょう。

## さいごに

Rubyワンライナーは非常に便利で、awkでできることはだいたいできます。
でも、何でもできるからといってすべてRubyでやるべきではありません。
（これはawkやperlにも言えることですが）

UNIXコマンドには、1つのことをうまくやる専用のコマンドが多数存在するので、専用コマンドでできることは任せたほうが良いです。

* 正規表現で文字列を絞り込むなら `grep`
* 文字列の置換なら `sed`
* ソートするなら `sort`

これらのコマンドは歴史が長く、チューニングされ尽くしています。
パフォーマンスの面で見ても、専用コマンドを使ったほうが良いです。

Rubyを使ってワンライナーで処理をする時は「Rubyでないと難しい」処理にのみ、Rubyを使うようにしましょう。