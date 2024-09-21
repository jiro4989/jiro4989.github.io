---
layout: default
title: "rpm パッケージを作成する GitHub Actions を作った"
date: 2020-07-14 09:00:00 +0900
categories: tech
---

# rpm パッケージを作成する GitHub Actions を作った

簡単に rpm パッケージを作れる GitHub Actions を作ってみた。

* Table of contents
{:toc}

## 成果物

* ソースコード: [jiro4989/build-rpm-action - GitHub](https://github.com/jiro4989/build-rpm-action)
* Marketplace: <https://github.com/marketplace/actions/build-rpm-action>

## 使い方

上のリポジトリの README に書いてあるとおり。
`package_root` で指定したパスにコマンドを配置するだけ。
`package_root` 配下のディレクトリ構造が、そのままインストール先になる。
こういう作りにしたのは、別記事で書いてる build-deb-action と同じ使い勝手にしたかったから。

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
          mkdir -p .rpmpkg/usr/bin
          mkdir -p .rpmpkg/usr/lib/testbin
          echo -e "echo hello" > .rpmpkg/usr/bin/testbin
          echo -e "echo hello2" > .rpmpkg/usr/bin/testbin2
          echo -e "a=1" > .rpmpkg/usr/lib/testbin/testbin.conf
          chmod +x .rpmpkg/usr/bin/*

      - uses: jiro4989/build-rpm-action@v1
        with:
          summary: 'testbin is a test script'
          package: testbin
          package_root: .rpmpkg
          maintainer: jiro4989
          version: '${{ steps.vars.outputs.tag  }}' # vX.X.X
          arch: 'x86_64'
          desc: 'test package'
```

## 使用例

nimjson という自作のツールで実際に使ってみた。

* <https://github.com/jiro4989/nimjson>
* [.github/workflows/release.yml](https://github.com/jiro4989/nimjson/blob/09d5b6ee0e765704c45dae12fe88edec2b732faf/.github/workflows/release.yml#L69-L86)

リリースを publish すると deb をリリースするようなワークフローにしている。
成果物は [Releases](https://github.com/jiro4989/nimjson/releases) にアップされる。
