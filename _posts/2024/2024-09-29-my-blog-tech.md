---
layout: default
title: "本ブログを支える技術"
date: 2024-09-29 14:58:24 +0900
categories: tech
---

# 本ブログを支える技術

ある程度ブログの体裁が固まったので、現時点で本ブログに使っている技術要素をまとめることにした。

* Table of contents
{:toc}

## 運用

ソースコードは <https://github.com/jiro4989/jiro4989.github.io> で管理している。

運用としては、以下の手順をしているだけ。

1. `scripts/ne_post.sh` を実行して記事を作成する
1. ファイルを編集する
1. `scripts/set_date.sh` を実行して index.md にリンクを埋め込み
1. `git push`

していることはたったこれだけだけれど、この中にいろんな技術要素がある。

## 技術要素

### GitHub Pages

まずこのサイトは [GitHub Pages](https://pages.github.com/) でホスティングしている。

GitHub Pages は GitHub が提供している静的ページを公開できるサービス。
無料で利用できる。

用途は主に 2 つある。

1. プロジェクト用のページ
1. ブログ

プロジェクト用のページというのは、例えば作ったライブラリの API ドキュメントのページ。
Nim や Go などでは、ドキュメンテーションコメントを記述することで、それらを API ドキュメントとして静的 HTML 出力する機能がある。
これらを GitHub Pages で公開すれば、ライブラリユーザに API 仕様を説明できる。

ブログは僕が本ブログを公開しているのがそれ。
GitHub Pages は静的サイトジェネレーターとして内部で [Jekyll](https://jekyllrb-ja.github.io/) を使用している。
ブログとして使うときに便利ないろんな変数が用意されている。

### Jekyll

前述のとおり、Jekyll は静的サイトジェネレータです。
Markdown で書いたものを HTML に変換してくれたり、Jekyll 固有の構文を使うことで、何らかの変数を HTML 内に埋め込んだりできる。

本ブログを作る上で Jekyll の機能を意識することはあまりなかったが、
唯一確認したのは[変数](https://jekyllrb-ja.github.io/docs/variables/)の説明です。
この中のサイト変数、ページ変数を使用して HTML に動的要素を埋め込んだ。

例えばページ変数を使ってこういうコードを書いた。
これはブログの各記事に、前後の記事への導線を配置するコード。
条件分岐を使って、前後要素があるときだけ `a` タグを作るようにしている。

```html
<p>
{% if page.previous %}<a href="{{ page.previous.url }}">前の記事 {{ page.previous.title }}</a>{% else %}前の記事{% endif %}
|
{% if page.next %}<a href="{{ page.next.url }}">次の記事 {{ page.next.title }}</a>{% else %}次の記事{% endif %}
</p>
```

この前後記事への導線は、使用しているテーマの [slate](https://github.com/pages-themes/slate) のデフォルトレイアウトには含まれていなかった。
別に動線も必要なかったけれど、簡単に仕込めそうだったので入れた。

### Markdown

わざわざ書く必要もない気がするけれど、当然 Markdown を使っている。
ブログ記事はすべて Markdown で書いている。

Markdown には方言があるので、書いた Markdown が描画されたときに正しく描画されないことがある。
GitHub においては [GitHub Flavored Markdown](https://github.github.com/gfm/) が使われている。
Markdown の書き方は[基本的な書き方とフォーマットの構文](https://docs.github.com/ja/get-started/writing-on-github/getting-started-with-writing-and-formatting-on-github/basic-writing-and-formatting-syntax)という記事で説明されている。

また、Markdown だけでなく静的サイトジェネレータの構文も使用できるため、
書いているコードが何の機能のコードなのか意識しなければならない。
例えば GitHub では以下の 2 行のコードで目次（Table of Contents）を自動生成できる。
このコードは、何の機能だろうか？

```
* Table of contents
{:toc}
```

これは Markdown の機能ではない。
Jekyll が使っている Markdown レンダラーの Kramdown の機能です。
[Kramdown は Jekyll のデフォルト Markdown レンダラー](https://jekyllrb-ja.github.io/docs/configuration/markdown/)です。
[Kramdown のドキュメント](https://kramdown.gettalong.org/converter/html.html)にて、`toc` について以下の記載がある。

> **Automatic “Table of Contents” Generation**
>
> kramdown supports the automatic generation of the table of contents of all
> headers that have an ID set.  Just assign the reference name “toc” to an
> ordered or unordered list by using an IAL and the list will be replaced with
> the actual table of contents, rendered as nested unordered lists if “toc”
> was applied to an unordered list or else as nested ordered lists.  All
> attributes applied to the original list will also be applied to the generated
> TOC list and it will get an ID of markdown-toc if no ID was set.
>
> When the auto\_ids option is set, all headers will appear in the table of
> contents as they all will have an ID.  Assign the class name “no\_toc” to a
> header to exclude it from the table of contents.
>
> Here is an example that generates a “Table of Contents” as an unordered list:

```md
# Contents header
{:.no_toc}

* A markdown unordered list which will be replaced with the ToC, excluding the "Contents header" from above
{:toc}

# H1 header

## H2 header
```

この説明から、toc が Kramdown の機能であることが分かる。
また、`no_toc` を設定することで H1 見出しを目次から除外する機能もあることが分かる。
使ったことはないが、おそらく使用できるのだろう。

とにかく `toc` は Markdown の機能ではなく Jekyll と Kramdown の機能です。
したがって、Jekyll を使っていない Markdown レンダリングをしているサイトに `toc` を書いても、目次を自動生成できない。

例として、おそらくはてなブログなどに `toc` を書いてもレンダリングされないだろう。
そもそも[はてなブログには `[:contents]` で目次を生成する機能がある](https://help.hatenablog.com/entry/markup/hatena/contents)のでなおさらそうだろう。

### HTML、CSS、JavaScript

これもわざわざ書く必要はない気がするけれど、当然 HTML、CSS、JavaScript を使っている。
slate テーマをそのまま使っていて、レイアウトに手を加える必要がなければ書く必要はない。
僕はレイアウトを調整したい箇所があったので、一部これらに手を加えている。

GitHub Pages でサイトを公開するだけであれば、Markdown の知識があるだけで一応ページを公開できる。
しかし Markdown しか書かなかったとしても、HTML 周辺の知識は一定把握しておいたほうが良い。
Markdown は最終的に HTML に変換されて出力される。
したがって、書いた Markdown がどのような HTML として出力されるかは、意識したほうがいい。

例えば、以下の Markdown がある。
これは画像ファイルを埋め込むための構文です。

```md
![](sushi.png)
```

これを HTML に変換すると、次のようになる。

```html
<p><img src="sushi.png" alt=""></p>
```

つまり Markdown における画像埋め込み構文の `[]` 内は、HTML の `img` タグにおける alt 属性です。

これを意識していないと何が起きるかというと、
例えば [WCAG（Web Content Accessibility Guidelines）](https://waic.jp/translations/WCAG22/)を守れないといったことが起きる。

<!-- textlint-disable -->

WCAG はウェブコンテンツをよりアクセシブルにするための広範囲に及ぶ推奨事項を網羅したガイドラインです。
このガイドラインに従うことで、全盲又はロービジョン、ろう又は難聴、運動制限、発話困難、光感受性発作及びこれらの組合せ、
並びに学習障害及び認知限界への一部の適応を含んだ、様々な障害のある人に対して、コンテンツをアクセシブルにすることができる。

<!-- textlint-enable -->

例えば WCAG 2.1 の「[達成基準 1.1.1 非テキストコンテンツ](https://waic.jp/translations/WCAG21/#non-text-content)」には、以下のガイドラインが示されている。

> 達成基準 1.1.1 非テキストコンテンツ
>
> （レベル A）
>
> 利用者に提示されるすべての非テキストコンテンツには、同等の目的を果たすテキストによる代替が提供されている。

<!-- textlint-disable -->

これに対する失敗例として[F65: 達成基準 1.1.1 の失敗例 － img 要素、area 要素、及び type "image" の input 要素の alt 属性又はテキストによる代替を省略している](https://waic.jp/translations/WCAG-TECHS/F65)がある。

<!-- textlint-enable -->

コード例は以下。

```html
<img src="../images/animal.jpg" />
```

つまり[非テキストコンテンツ](https://waic.jp/translations/WCAG21/#dfn-non-text-content)には、適切な代替テキストを設定することが求められる。
これを意識するには `img` タグにおける alt 属性を知っている必要があり、この HTML 要素と Markdown の関係性を知っていなければならない。

また別の例を挙げる。
以下の Markdown がある。

```md
台風 999 号が南から接近しています。
3 日後には本州に上陸する見通しです。

[Read More ](ff.html)
```

これは HTML だと以下になる。

```html
<p>台風 999 号が南から接近しています。
3 日後には本州に上陸する見通しです。</p>
<p><a href="ff.html">Read More </a></p>
```

この例だと WCAG 2.1 の[達成基準 2.4.4 リンクの目的 (コンテキスト内)](https://waic.jp/translations/WCAG21/#link-purpose-in-context)に違反する。

> 達成基準 2.4.4 リンクの目的 (コンテキスト内)
>
> (レベル A)
>
> それぞれのリンクの目的が、リンクのテキスト単独で、又はリンクのテキストとプログラムによる解釈が可能なリンクのコンテキストから判断できる。
> ただし、リンクの目的がほとんどの利用者にとって曖昧な場合は除く。

この Read More というリンクは、それ単体でリンクの目的が解釈できない。
段落としても分かれているため、プログラムとしてもリンクの関連を理解できない。

失敗例は[F63: 達成基準 2.4.4 の失敗例 － リンクと関係のないコンテンツにのみ、リンクの文脈を提供している](https://waic.jp/translations/WCAG-TECHS/F63.html)としてまとまっている。

> あるニュースサービスでは記事の冒頭のいくつかの文を一つの段落に入れている。
> その次の段落には「Read More...」というリンクが置かれている。
> そのリンクは導入文と同じ段落にないので、利用者はそのリンクが何についての続きを読むのかを容易に見つけることができない。

```html
<p>A British businessman has racked up 2 million flyer miles and plans to 
travel on the world's first commercial tourism flights to space.</p>

<p><a href="ff.html">Read More...</a></p>
```

ではどう直せばいいかというと、こうすればいい。
間の空白行を削除する。

```md
台風 999 号が南から接近しています。
3 日後には本州に上陸する見通しです。
[Read More ](tenkiyoho.html)
```

これが HTML に変換されると、同じ段落(`p`)になる。
そもそも Read More ってリンクの貼り方をやめたほうが良いと思うが、それはまた別の話。

```html
<p>台風 999 号が南から接近しています。
3 日後には本州に上陸する見通しです。
<a href="tenkiyoho.html">Read More </a></p>
```

これは Markdown では空白行が段落の区切れ目として解釈されるから、このようなことになる。
先の例では、間に空白行が存在したため、文章とリンクはそれぞれ別の段落になっていた。

このように、Markdown がどのように HTML に変換されるかを意識していないと、困る場合がある。
WCAG を守るかどうかに限らず、読みやすい文章を意識するならば Markdown しか書かなくても HTML 周りは知っていたほうがいい。

### Go 言語

一部の手作業を簡略化するためにツールを Go で書いている。
Go にしたのは僕が使い慣れているからというだけで、別に Ruby や Python でも良かった。
それなりに複雑なコードで、テストコードも書きたかったので、静的型付け言語の Go で書いた。

コードとしては scripts ディレクトリ配下の embed\_links と generate\_markdown\_links がそれです。

### テンプレートエンジン

前述の Jekyll のテンプレートコードもそうだが、Go のツールでもテンプレート([text/template](https://pkg.go.dev/text/template))を使っている。
やっている処理は、Markdown の投稿記事リンクの生成。
index.md の投稿記事セクションに埋め込む用途で使っている。
以下のようなリンクを生成する。

```md
* 2024-07-05 技術 [setup-nim-action を v2 にバージョンアップした](/tech/2024/07/05/setup-nim-action-v2.html)
```

この程度のリンク生成であれば、わざわざ text/template のようなテンプレートエンジンを使う必要はない。
しかし構造化された文章を繰り返し生成するといったユースケースは、テンプレートエンジンと非常に相性が良いので採用した。

テンプレートエンジンの使用経験としては、圧倒的に Go のテンプレートエンジンのが使い慣れている。
他には Ansible で使われている Jinja2 も使ったことがある。
そして Jekyll のテンプレートエンジンは全く触ったことがなかった。
それでも、多少ドキュメントを読めば少しのテンプレート程度ならすぐに埋め込める。

テンプレートエンジンはそれぞれ構文が多少異なるが、考え方はどれも似通っている。
なにか 1 つをある程度使い込んだ経験があれば、他のテンプレートエンジンもスムーズに使っていけると思う。

### ユニットテスト

単純なシェルスクリプト程度ならユニットテストは書かないが、それなりに複雑な処理を書いたときは必ずユニットテストを実装する。
本ブログにおいては、embed\_links と generate\_markdown\_links でユニットテストを実装した。

Go はユニットテストの機能が言語自体に組み込まれているため、非常にテストコードが書きやすい。
ただし、ユニットテスト用の関数などはあまり充実していないため、サードパーティのライブラリを使うケースがほとんどだと思う。
Go では [testify](https://github.com/stretchr/testify) がおそらく最もメジャーなテストライブラリだろう。
仕事でもプライベートでもよく使っている。

Go 言語でユニットテストのコードを書くときは、テーブルドリブンテストを書くことが多い。
テーブルドリブンテストは [Go の Wiki でも専用ページが設けられている](https://go.dev/wiki/TableDrivenTests)ため、Go においてはメジャーな手法だ。

#### テストコードを書くときのルール

コード的には以下のコードを書いた。
他にもいくつかテストケースを書いたが、概ね書き方は同じ。

```go
func TestReadAttrLine(t *testing.T) {
    tests := []struct {
        desc    string
        path    string
        attr    string
        want    string
        wantErr bool
    }{
        {
            desc:    "正常系: 最初の layout を読み取る",
            path:    "sample1.md",
            attr:    "layout:",
            want:    "default",
            wantErr: false,
        },
        {
            desc:    "異常系: 該当する attr が見つからない",
            path:    "sample1.md",
            attr:    "pohe:",
            want:    "",
            wantErr: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.desc, func(t *testing.T) {
            a := assert.New(t)

            path := filepath.Join("testdata", tt.path)
            got, err := readAttrLine(path, tt.attr)
            if tt.wantErr {
                a.Error(err)
                a.Empty(got)
                return
            }
            a.Equal(tt.want, got)
            a.NoError(err)
        })
    }
}
```

僕がテストコードを書くときは、いくつか一貫したルールを設けて書いている。
このルールはあくまでも僕個人が意識しているだけのものであり、何かのガイドラインに従ったものではない。

1. テストケース構造体は無名構造体にして、テスト関数内で定義する
1. テストケースの説明(desc)を書く
1. テストケースの説明には正常系と異常系のどちらなのかを書く
1. テストケースの順序は正常系を先に書く
1. テストケースの期待値変数は `want` にする
1. テストケースの期待値エラー変数は `wantErr (bool)` にする
1. テスト対象の関数の戻り値変数名は `got` にする

テストケース構造体を無名構造体にしているのは、その型定義を他のテストケースで使うことがほぼなくて、名前をつける意味がないため。

テストケースの説明は、他の人がテストコードを見たときに、テストの目的を理解できるようにするため。
また、説明を補足するために正常系、異常系を明示する。

正常系を先に書くのは、テストコードを他の人が読んだとき、まず上からテストコードを読むだろうから。
コードを理解するなら、まず正常系を読んで異常系を読むほうが理解しやすいだろう、との配慮でそうしている。

変数名の統一は、変数名から目的をパッと理解できるようにするため。
あと grep しやすいから。

#### テストする前提でコードを書く

こういう関数を書いたが、引数として `io.Reader` インタフェースを渡すようにしている。

```go
// readStdin は標準入力を文字列の配列として返す。
func readStdin(r io.Reader) ([]string, error) {
    ret := make([]string, 0)
    sc := bufio.NewScanner(r)
    for sc.Scan() {
        line := strings.TrimSpace(sc.Text())
        ret = append(ret, line)
    }
    if err := sc.Err(); err != nil {
        return nil, err
    }
    return ret, nil
}
```

ツールとしては標準入力からしかデータを渡さないので、別にインタフェースしなくて良いといえば良い。
`bufio.NewScanner` に `os.Stdin` を直接渡してもいい。
でも、あえてそうしていない。
理由はテストコードを書きやすくするため。

標準入出力や、ファイル IO みたいな外部とやりとりするコードを直接書いてしまうと、テストコードが書きづらくなる。
しかしインタフェースを渡すように実装してあると、外部を意識せずにテストがかけて保守性が上がる。

たとえば先の標準入力部分のテストは以下の実装になっている。

```go
a := assert.New(t)

r := strings.NewReader(tt.text)
got, err := readStdin(r)
if tt.wantErr {
    a.Error(err)
    a.Nil(got)
    return
}

a.Equal(tt.want, got)
a.NoError(err)
```

標準入力をあれこれいじってテキストを流し込んだりとかする必要がない。
ただの文字列を strings.Reader にして、そのまま渡すことができる。
非常にシンプルなコードになった。

標準出力を扱うコードを書く場合も `fmt.Println` で標準出力に書くのではなく
`fmt.Fprintln` に `io.Writer` を渡すように書いたほうがいい。
そうすれば、以下のように `bytes.Buffer` を使って書き込まれたテキストを簡単に取り出してテスト結果を比較できる。

```go
a := assert.New(t)

w := &bytes.Buffer{}
err := generateLinks(tt.i, w)
if tt.wantErr {
    a.Error(err)
    return
}

got := w.String()
a.Equal(tt.want, got)
```

### npm

後述する textlint を管理するために npm を使っている。
これは Node.js のパッケージマネージャだ。

Node.js にはあまり明るくないが、最低限の仕組みは知っている。

### textlint

[textlint](https://github.com/textlint/textlint) は Node.js 製の文章構成ツールです。
日本語をサポートしていて非常に重宝している。

一文が長すぎたり、「が」が何度も登場して文章が複雑になっているのを検出してくれる。
これにより一文がシンプルで短くなって、読む人の負担を減らしてくれる。

文章の執筆には vim を使っているが textlint が書いている文章を自動で指摘してくれるため容易に修正できる。
こういったエディタとの連携は各エディタがやっていて、VSCode にも同様のものがある。
長い文章を執筆する際はエディタと連携するのをオススメする。

### GitHub Actions

[GitHub Actions](https://docs.github.com/ja/actions) は GitHub に組み込まれている CI/CD 基盤です。

Push、Pull Request などをトリガーに CI を走らせたり、特定のファイルが変更されたときだけ CI を走らせたりできる。
本ブログでは textlint とツールのユニットテストを GitHub Actions でテストしている。

テストの条件はファイルの変更があったときだけ実行するようにしている。
Markdown 記事が編集されたときだけ textlint を実行する。
scripts 配下が変更されたときだけユニットテストを実行する、といった具合。

### dependabot

[dependabot](https://docs.github.com/ja/code-security/dependabot/working-with-dependabot) は
GitHub に組み込まれたパッケージ自動更新の仕組みです。

新しいバージョンのパッケージが公開されたら自動で Pull Request を作ってくれる。
先の textlint なども、定期的にアップデートが入るため、最新に追従する目的で導入している。
自動で Pull Request が作られて、CI がパスしたら自動でマージするようにしているため、運用の負担はほとんどない。
