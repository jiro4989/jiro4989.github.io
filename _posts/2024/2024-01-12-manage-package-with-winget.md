---
layout: default
title: "winget でパッケージを管理する"
date: 2024-01-12 09:00:00 +0900
categories: tech
categoriesJP: 技術
---

# winget でパッケージを管理する

2024/1/12 にデスクトップ PC を新しくした。
新しい PC の開発環境をセットアップして、いつものように scoop をセットアップしたら、変な挙動をしていた。

```ps1
Set-ExecutionPolicy -ExecutionPolicy RemoteSigned -Scope CurrentUser
Invoke-RestMethod -Uri https://get.scoop.sh | Invoke-Expression
scoop install git
```

これを実行すると、git のインストール前にインストールされる 7zip のインストールで失敗していた。

エラーメッセージを読むに、git コマンドを使おうとしているが git コマンドが存在しなくてエラーになっていた。
そこで、scoop 周りの情報を調べていたら、以下の記事を見つけた。

* [scoopシステム崩壊の序曲](https://zenn.dev/zetamatta/scraps/b21750b7ac7c06)

どうやら scoop のメンテナンスが停止してしまっている可能性があるらしい。
これはどうしたものか。scoop から chocolatey に戻るべきか？と思っていたが winget の存在を思い出した。

* [winget ツールを使用したアプリケーションのインストールと管理](https://learn.microsoft.com/ja-jp/windows/package-manager/winget/)

これは Microsoft 公式がサポートするツールなので、信頼できるだろう。
ということで、今までずっと使っていた scoop を使うのをやめて winget を使ってみる。

* Table of contents
{:toc}

## winget のインストール

[winget ツールを使用したアプリケーションのインストールと管理](https://learn.microsoft.com/ja-jp/windows/package-manager/winget/)にかかれているコマンドを実行してインストールする。

```ps1
Add-AppxPackage -RegisterByFamilyName -MainPackage Microsoft.DesktopAppInstaller_8wekyb3d8bbwe
```

そして `winget search vscode` すると、**応答がない**。
調査したところ、以下の記事を発見した。

* [wingetが応答しないときの対処方法](https://zenn.dev/kawamasato/articles/11c4477e101374)

どうやらこ以下のバイナリっぽい。

* <https://github.com/microsoft/winget-cli/releases>

上記記事を参考に winget を更新すると、winget のバージョンが上がった。

```ps1
$ winget -v
v1.6.3482
```

これでようやく動作するようになった。

## winget を使ってみる

早速 git をインストールしてみる。

```ps1
$ winget search git
..
Git   Git.Git  2.43.0  winget

$ winget install Git
複数のパッケージが入力条件に一致しました。入力内容を修正してください。
名前   ID           ソース
---------------------------
My Git 9NLVK2SL2SSP msstore
Git    Git.Git      winget
```

どうやら `git` だけだと複数のパッケージがヒットするらしく、`Git.Git` と入力しないとだめらしい。
以下のように実行することでインストールできた。

```ps1
winget install Git.Git
```
