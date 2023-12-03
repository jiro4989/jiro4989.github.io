---
layout: default
title: "Bloodborne のビルドシミュレータを Next.js で作った"
date: 2023-12-03 20:48:20 +0900
categories: tech
---

# Bloodborne のビルドシミュレータを Next.js で作った

最近 Bloodborne をずっと遊んでる。
フロムゲーは今までまったくやったことなくて、エルデンリングが初めてのフロムゲーだった。
よって Bloodborne は、僕がプレイする2つ目のフロムゲーということになる。

Bloodborne はエルデンリングと違ってステータスの振り直しが不可能だった。
なので、新しいビルドを試すには、最初からやり直すか、ひたすらレベルを上げ続けるかのどっちかしかない。

ステータスを振り直せない以上、事前にステ振り方針を決めておく必要がある。
そこで Bloodborne のビルドシミュレータを探してたんだけれど、ビルドを共有できるシミュレータが存在しなかった。

既存のシミュレータは以下。

* [Bloodborne（ブラッドボーン）パラメータシミュレータ](https://game.kozuyu.com/bloodborne/)
* [Bloodborne](https://mugenmonkey.com/bloodborne)

おそらく海外製のシミュレータの方は共有できるっぽかったけれど、機能過多すぎる感もあって、もっとシンプルなものがほしかった。
[ELDEN RING ビルドシミュレータ](https://8bitdesign.dev/ja/elden-ring/)などが理想だったけれど、これに近いものはなさそうだった。

ないのなら、自分で作ってしまえば良い。
ということで、Bloodborne のビルドシミュレータを作ってみた。

* Table of contents
{:toc}

## サイト

以下の URL で公開している。URL の通り GitHub Pages でホスティングしている。

* <https://jiro4989.github.io/bloodborne-build-simulator/>

![Bloodborne ビルドシミュレータのトップページ](https://github.com/jiro4989/bloodborne-build-simulator/raw/main/docs/toppage.png)

## 技術要素

Next.js と TypeScript で作ってみた。
CSS は [Tailwind CSS](https://tailwindcss.com/) を使っている。

Next.js を選択した理由は、主に以下の2つ。

1. React.js は使ったことがあるので、それのフレームワークの Next.js が気になっていた
1. 数値変動で色々要素が変化する想定だったので、Next.js と相性が良さそうだった

この程度の規模のアプリなら Next.js を使うのは大掛かりすぎるけれど、技術的に使ってみたかったので採用した。

CSS に Tailwind を採用したのは `npx create-next-app` のプロンプトに表示されていたから。
まぁ、ついでに有効化したというだけ。
そもそもこのフレームワークの存在を知らなかったので、最初はCSSを全部手書きしようかと思っていた。

## 重視した箇所

今回アプリを作る上で重視したのは、以下の2点。

1. GitHub Pagesでホスティングできること
1. URLで共有できること
1. 複雑にしすぎない

クエリストリングを使えば静的サイトでも十分動作するし、仕組みも単純になる。
そしてブックマークできる。

海外のシミュレータだと、もっと大量のパラメータを調整できるようだったが、そこまでにする気はなかった。
HPとスタミナとレベルの計算だけあれば十分だったので、機能は最小限にした。

## 制作時間

Togglで記録していた感じだと約7時間ほど。
まぁ途中動画見たりして遊んでたので、もう少し短いかもしれない。

## ハマった箇所

今回初めて Next.js を触ったけれど、React.jsは [jira\_issue\_url\_generator](https://github.com/jiro4989/jira_issue_url_generator) で触ったことがあった。
が、Next.js のお作法とかは知らなかったので、ちょいちょい詰まった箇所がある。

### 静的ファイル出力するには `next.config.js` に設定が必要

以前は `next export` コマンドでファイル出力できてたらしいけれど、現在は異なるようだった。
`next export` コマンドを実行すると、以下のエラーメッセージが表示される。

```bash
⟩ npm run export

> bloodborne-build-simulator@0.1.0 ex
> next export

 ⨯ The "next export" command has been removed in favor of "output: export" in next.config.js. Learn more: https://nextjs.org/docs/advanced-features/static-html-export
```

そして、そちらの URL 先の記事を読むに、設定ファイルで制御するように変わったらしい。
`next.config.js` に以下の設定を書く。

```js
const nextConfig = {
  output: 'export', // 静的ファイル出力を有効化
}

module.exports = nextConfig
```

上記設定をすれば `next build` で、ファイルが `out` ディレクトリに出力される。

### 環境ごとに URL を切り替える

ローカル環境では localhost で、GitHub Pages 上では Pages のドメインを使うように切り替える必要がある。
これには `.env.${ENV}` ファイルを使うことで実現できる。

まず、以下のファイルを作る。

| 環境 | ファイル |
| --- | --- |
| 本番 | .env.production |
| 開発 | .env.development |

そして、それぞれファイルの中身は以下のとおり。

```.env.production
// .env.production
NEXT_PUBLIC_BASE_URL=https://jiro4989.github.io/bloodborne-build-simulator/
```

```.env.development
// .env.development
NEXT_PUBLIC_BASE_URL=http://localhost:1323
```

これらがそのまま環境変数として使用可能になるが、
クライアントサイドで使用する環境変数には `NEXT_PUBLIC_` というプレフィックスが必須である。

参考: [Bundling Environment Variables for the Browser](https://nextjs.org/docs/pages/building-your-application/configuring/environment-variables#bundling-environment-variables-for-the-browser)

また、これらのファイルが使用されるタイミングも決まっている。

* `next dev` では `.env.development`
* `next start` では `.env.production`
* `next build` では `.env.production`
  * これは Default Environment Variables のセクションには明記されてない
  * Bundling Environment Variables for the Browser のセクションに書かれている

参考: [Default Environment Variables](https://nextjs.org/docs/pages/building-your-application/configuring/environment-variables#default-environment-variables)

実際にコマンドを実行してみると、読み込んでいるファイルが表示される。

```bash
⟩ npm run build

> bloodborne-build-simulator@0.1.0 build
> next build

   ▲ Next.js 14.0.3
   - Environments: .env.production // <- ここ
```

### アセットパスの切り替え

これは GitHub Pages 固有の問題になる。

GitHub Pages で静的サイトをホスティングする際、各プロジェクト用の Pages のベースURL は以下のようになる。

* `https://<Username>.github.io/<Repository>/`

これに対して、Next.js が生成した静的ファイルのパスは、以下のようになる。

* `/_next/static/chunks/main-app-*.js`

つまり、静的ファイルのパスの前に、リポジトリ名のパスが必要になる。

ローカル環境であれば、以下のパスで問題ない。

* `http://localhost:1323/_next/static/chunks/main-app-*.js`

が、GitHub Pages 上で動かすなら、以下のようにしなければならない。

* `https://<UserName>.github.io/<Repository>/_next/static/chunks/main-app-*.js`

よって、環境に応じてアセットパスの有無を切り替えなければならない。
これも、前述の `.env.${ENV}` ファイルで解決できる。

以下のようにする。

```.env.production
// .env.production
NEXT_PUBLIC_BASE_URL=https://jiro4989.github.io/bloodborne-build-simulator/
ASSET_PREFIX=/bloodborne-build-simulator
```

```.env.development
// .env.development
NEXT_PUBLIC_BASE_URL=http://localhost:1323
ASSET_PREFIX=
```

そして、next.config.js に以下の設定を追加する。

```js
const nextConfig = {
  assetPrefix: process.env.ASSET_PREFIX,
}

module.exports = nextConfig
```

前述の `NEXT_PUBLIC_BASE_URL` と違って、 `ASSET_PREFIX` はサーバサイドで使う環境変数である。
したがって `NEXT_PUBLIC_` プレフィックスは不要である。

### GitHub Pages へのデプロイで権限不足

GitHub Pages にデプロイするのは、いつもどおり GitHub Actions からやることにした。
`jira_issue_url_generator` の設定をコピーしてきて調整して、以下の記述でデプロイするようにした。

```yaml
---
name: pages

"on":
  push:
    tags:
      - 'v*'

jobs:
  pages:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3

      - uses: actions/setup-node@v4
        with:
          node-version: '20'

      - name: Install
        run: npm ci

      - name: Build
        run: npm run build

      - name: Deploy pages
        uses: peaceiris/actions-gh-pages@v3
        with:
          github_token: ${{ secrets.GITHUB_TOKEN }}
          publish_dir: ./out
```

しかし、権限エラーが発生していた。

```
Push the commit or tag
  /usr/bin/git push origin gh-pages
  remote: Permission to jiro4989/bloodborne-build-simulator.git denied to github-actions[bot].
  fatal: unable to access 'https://github.com/jiro4989/bloodborne-build-simulator.git/': The requested URL returned error: 403
  Error: Action failed with "The process '/usr/bin/git' failed with exit code 128"
```

どうやらリポジトリの CI の権限が、細分化されたようで、明示的に許可しないと権限不足になるらしい。
権限変更するには、以下の手順で変更できる。

1. リポジトリの Settings 画面を表示する
1. 「General -> Actions -> General」を表示する
1. 「Workflow permissions」を「Read and write permissions」に変更する

これで失敗したジョブを再実行すると、成功した。

## 感想

今回初めて Next.js を触ったが、React.js と感触がほとんど変わらなくておどろいた。
書き方などはほとんど React.js の書き方がそのまま使える。
ハマったのはやはり Next.js 固有の設定周りだった。
このあたりは仕方ないだろう。

ただこの程度の規模のプロジェクトだと Next.js を使った恩恵はあまり感じられなかった。
感覚的には React.js と何が違うのか分からないくらい、差が分からなかった。

もうすこし大規模なプロジェクトになれば、また違ってくるのだろうか？

Tailwind CSS については、簡単に触った程度となってしまった。
https://tailwindcss.com/docs/margin などのドキュメントを見ながら、
提供されているクラスを探して適用するだけで使えるので、CSSに疎い僕には便利だった。

ただし、CSSクラスの当て方を雑にやってしまったので、ソースコードはかなり汚い。
きれいに HTML 要素にクラスを当てる方法については知見がないので、このあたりは自分の課題だと感じた。
