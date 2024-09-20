---
layout: default
title: "Node プロジェクトで、ビルドプロセスでのみ必要なパッケージを dependencies で管理するのは誤り"
date: 2023-12-06 09:00:00 +0900
categories: tech
---

# Node プロジェクトで、ビルドプロセスでのみ必要なパッケージを dependencies で管理するのは誤り

本番環境向けのビルド用 CI を整備していたときに dependencies と devDependencies の使い分けに悩んだ。
この記事はその調査結果をまとめたもの。

* Table of contents
{:toc}

## 導入

仕事で CI ジョブにかかっている時間の削減のために、`npm ci`周りを見直していた。
`npm ci`は package-lock.json に書かれてるファイルを取得して node\_modules 配下にインストールするコマンド。
つまり依存パッケージを取得するためのコマンドである。

`npm ci`ではオプションで `--omit=dev` が指定できる。
これは `devDependencies`をインストール対象から除外するオプションである。

devDependencies は、開発時にのみ必要なパッケージなどを指定するパラメータである。
したがって `--omit=dev`を使えば、本番向けビルドに不要なパッケージ(jest とか eslint みたいなテストや静的解析でしか使わないツール郡)のインストール時間を削減できる。

しかしながら、devDependencies に、本番環境での動作には必要ないが、ビルドプロセスで必要なパッケージがいくつか含まれていた。
typescript とか webpack とかそういう類のものである。

TypeScript でコードを書いていても、それをそのまま本番環境で動作させることはできない。
一度 Javascript にトランスコンパイルしないと、当然ブラウザ上では動作しない。
しかし、TypeScript はブラウザ上で必要ない。
こういったパッケージは devDependencies ではなく dependencies で管理するべきだろうか？

## 他の人は dependencies と devDependencies をどう使い分けている？

同様のことを考えた方が他にもいた。
例えば以下の記事では「Webアプリケーション開発とライブラリ開発では devDependencies に含めるパッケージは異なる」といった主張をされている。

* [dependencies と devDependencies の使い分け - Panda Noir](https://www.pandanoir.info/entry/2020/07/08/193000)

> 大抵、本番環境ではリポジトリをクローンしてきてビルドをするだけです。
> そのため、ESLint や Prettier は本番環境で使わないものは devDependencies に書きます。
> 反対に、本番環境でも使いたいものは dependencies に書きます。
>
> * dependencies に書くパッケージ
>   * ビルドに必要なパッケージ(webpack、TypeScript、Babelなど)
>   * 使用するライブラリ・フレームワーク(React や Vue)
> * devDependencies に書くパッケージ
>   * Linter・Formatter(ESLint、Prettier 関連のパッケージ)
>   * テストフレームワーク(Jest など)

しかしながら、上記の記事には情報ソースがなかった。
他にも似たような主張をしている記事がいくつかある。

* [【package.json】dependencies, devDependencies の使い分けを考える](https://qiita.com/karur4n/items/3d9d28f6f21c3533020d)
* [/shokai/dependenciesにbabelを入れる](https://scrapbox.io/shokai/dependencies%E3%81%ABbabel%E3%82%92%E5%85%A5%E3%82%8C%E3%82%8B)

しかし、いずれも情報ソースが明記されていなかった。
僕も同様の考えを持っていたが、同じく考えを裏付ける根拠を持っていなかった。

公式ドキュメント以外の Web 記事を見る時はなるべく情報ソースが明記されているか確認するようにしている。
残念ながら上記の記事には主張を裏付ける情報ソースが記載されていなかったため、情報ソースを探すことにした。

## 公式ドキュメント

npm 公式のドキュメントで関連が深いのは以下の 2 箇所。

* [package.json dependencies](https://docs.npmjs.com/cli/v10/configuring-npm/package-json#dependencies)
* [package.json devDependencies](https://docs.npmjs.com/cli/v10/configuring-npm/package-json#dependencies)

まず、dependencies には次のように書かれている。

> Please do not put test harnesses or transpilers or other "development" time
> tools in your dependencies object.  See devDependencies, below.
>
> テスト ハーネス、トランスパイラー、またはその他の「開発」時ツールを
> dependencies オブジェクトに含めないでください。以下の devDependency を参照してください。

トランスパイラが明記されている。
TypeScript はトランスパイラなので、devDependencies に含めるべき、ということになる。

そして devDependencies では次のように書かれている。

> If someone is planning on downloading and using your module in their program,
> then they probably don't want or need to download and build the external test
> or documentation framework that you use. In this case, it's best to map
> these additional items in a devDependencies object.
>
> 誰かがあなたのモジュールをダウンロードして自分のプログラムで使用することを計画している場合、
> その人はおそらく、あなたが使用する外部のテスト フレームワークやドキュメント フレームワークをダウンロードして構築することを望まないか、またはその必要はありません。
> この場合、これらの追加項目を devDependency オブジェクトにマップするのが最善です。

これはその通りで、例えば prettier をインストールして jest までついてきたりしたら嫌だ。
しかし Web アプリケーションみたいに、他の Node プロジェクトから参照されないプロジェクトではどうしたらよいかが書かれていない。

これらの説明だけだと「ライブラリやツールとして配布する場合のみを想定して説明が書かれているのだろうか？」と感じてしまう。
ビルドして圧縮した JS や CSS だけを CDN で配布するようなケースにおいては、トランスパイラや圧縮ツールは本番向けビルドに必要である。
これらを dependencies に含めるのは果たして正しいのか？

CI の処理時間を早めるために、ビルドプロセスで必要なツールは dependencies に含めて、それ以外は devDependencies に含めるといった使い分けをしたかったのが本音である。
しかし、正しい使い方を明確にしておきたい。

## npm の人の説明

仕事で node プロジェクトに触れることはそれなりにある。
しかし、僕は Web フロントエンド事情に明るくない。
npm の公式ドキュメントも読み慣れていないので、もしかしたら公式ドキュメントのどこかに記載があるのかもしれない。
探した限りではそれらの記載を見つけられなかった。
npm のドキュメントリポジトリの issues や Pull Request も探したが、今回調べていた件に近しい議論は見つからなかった。

そこで npm のドキュメントリポジトリに issues を立てて「プロジェクト形態に合わせた dependencies と devDependencies の説明が追加で必要では」という起案をした。

* <https://github.com/npm/documentation/issues/862>

何度かやりとりをした結果、結論としては以下のとおり。

> Yes, that's the wrong usage. Some have asked for more granular categories
> like "test deps" and "build deps" - but all of these are basically always dev
> deps. dependencies is the place where everything that's conceptually needed
> **at runtime** goes, and devDependencies is where everything else goes,
> whether an application or a package.
>
> はい、それは間違った使い方です。
> 「test deps」や「build deps」などのより詳細なカテゴリを求める人もいますが、
> これらはすべて基本的に常に dev deps です。dependencies は**実行時に**概念的に
> 必要なものがすべて置かれる場所であり、devDependency はアプリケーションであれパッケージであれ、
> その他すべてが置かれる場所です。

したがって、**ビルドプロセスで必要となるパッケージであったとしても、実行時に不要なパッケージは devDependencies で管理するのが npm 公式の考える正しい使い方**である。

例を上げると、以下のような使い分けになる。

* ライブラリとして jQuery がフロントエンドで必要なケースなら、dependencies で管理する
* TypeScript はフロントエンドで実行時に直接使うことはないので、devDependencies で管理する
* webpack もフロントエンドで実行時に直接使うことはないので、devDependencies で管理する

あくまでも**実行時に**必要なパッケージのみを管理するのに dependencies を使うべき。

ただし、test deps, build deps といった話が出ているとおり、やはり僕と同様の意見を持っている人はいたらしい。
将来 npm がどうなるかは分からないが、少なくとも今日（2023/12/07）時点ではこの方針のようだ。

## まとめ

以上の情報をまとめると次の通り。

* 実行時に必要なパッケージを管理するために dependencies を使用する
* TypeScript や webpack などの実行時に不要なパッケージは、たとえビルドプロセスで必要だとしても devDependencies で管理するべき
* したがって、ビルドプロセスに不要なパッケージを除外する（`npm ci --omit=dev`）目的で、
  ビルドに関連するパッケージを dependencies で管理しようとするのは不適切である

「`--omit=dev`で除外すればテストでしか使ってないパッケージのインストール時間削減できるじゃん！」と思って調査し始めたわけだが、結論ダメだと分かった。
少なくとも npm の人はそういった使い方を想定していないようなので、やるなら自己責任ってことになる。
そのため僕は素直に `--omit=dev` は使わないことにした。
原理主義っぽい主張かもしれないが、プロジェクトにそこまで支障がない間は素直に公式の方針に従うのが無難だろう。

`devDependencies` のインストールだけで 10 分や 20 分かかるようになったら、リスクを承知で `--omit=dev` してもいいかもしれない。
まぁそんな状態になったら devDependencies の整理したほうが良いと思うが。

以上。
