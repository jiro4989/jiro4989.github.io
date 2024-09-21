---
layout: default
title: "debian パッケージを作成する GitHub Actions を作った"
date: 2020-07-10 09:00:00 +0900
categories: tech
---

# debian パッケージを作成する GitHub Actions を作った

簡単に debian パッケージを作れる GitHub Actions を作ってみた。

* Table of contents
{:toc}

## 成果物

* ソースコード: [jiro4989/build-deb-action - GitHub](https://github.com/jiro4989/build-deb-action)
* Marketplace: <https://github.com/marketplace/actions/build-deb-action>

## 使い方

上のリポジトリのREADMEに書いてあるとおり。
`package_root` で指定したパスにコマンドを配置するだけ。
`package_root` 配下のディレクトリ構造が、そのままインストール先になる。

```yaml
name: build

on:
  push:
    tags:
      - 'v*.*.*'

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2

      - name: Set tag
        id: vars
        run: echo ::set-output name=tag::${GITHUB_REF:10}

      - name: create sample script
        run: |
          mkdir -p .debpkg/usr/bin
          mkdir -p .debpkg/usr/lib/samplescript
          echo -e "echo sample" > .debpkg/usr/bin/samplescript
          chmod +x .debpkg/usr/bin/samplescript
          echo -e "a=1" > .debpkg/usr/lib/samplescript/samplescript.conf

      - uses: jiro4989/build-deb-action@v2
        with:
          package: samplescript
          package_root: .debpkg
          maintainer: your_name
          version: ${{ steps.vars.outputs.tag  }} # vX.X.X
          arch: 'amd64'
          desc: 'this is sample package.'
```

## 使用例

nimjson という自作のツールで実際に使ってみた。

* <https://github.com/jiro4989/nimjson>
* [.github/workflows/release.yml](https://github.com/jiro4989/nimjson/blob/09d5b6ee0e765704c45dae12fe88edec2b732faf/.github/workflows/release.yml#L69-L86)

リリースを publish すると deb をリリースするようなワークフローにしている。
成果物は [Releases](https://github.com/jiro4989/nimjson/releases) にアップされる。

## ポリシー

簡単に使用できて、使い方がわかりやすいものを目指した。

最初、僕も debian パッケージを作ろうとして、既存の GitHub Actions も探したけれど、いまいち使い方が分からないものばかりだった。
なので、直感的に使えるもので、見れば使い方が分かる、ってくらい単純なものとして build-deb-action を作成した。
ディレクトリ作って、その中に必要な実行可能ファイルを詰め込むだけで良いので、かなり単純。

実行可能ファイル 1 つ配布するだけの deb パッケージであれば、誰でも簡単に使えると思う。
preinst や postinst とかやりたくなったら、ちょっと工夫が必要だけどね。

以上。
