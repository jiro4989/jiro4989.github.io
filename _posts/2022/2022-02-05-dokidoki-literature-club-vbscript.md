---
layout: default
title: "ドキドキ文芸部プラス！のシークレット 10 を取得するための VBScript を書いた"
date: 2022-02-05 09:00:00 +0900
categories: tech
categoriesJP: 技術
---

# ドキドキ文芸部プラス！のシークレット 10 を取得するための VBScript を書いた

ドキドキ文系部プラス！の実績を全解除したけれど、
そのうち一番苦戦したシークレット 10 を楽に取得するために書いた VBScript についてまとめる。

シークレット 10 はめちゃくちゃ低い確率で 2 周目のメインメニュー（サユリがぶっ壊れて表示される画面）を表示した時に取得できる。
10 回以上再表示を繰り返しても取得できなくて、めんどくさくなったのでスクリプトを書いて取得した。

* Table of contents
{:toc}

## 注意点

VBScript を使ってキーボード入力を自動化したわけだけれど、
VBScript は Windows でしか使えないので、Linux や Mac では動かない。

僕が動作確認したのも Windows 11 なので、それ以外の環境は別の方法を自分で考える必要がある。
また、万が一変な動作しても自己責任で。

## 使い方

スクリプトは以下。`dokidoki_literature_club.vbs` という名前で保存する。

```vbs
'
' ドキドキ文芸部2周目のメインメニュー画面を表示した直後に実行する。
'
' 実行方法:
'   wscript.exe dokidoki_literature_club.vbs
'
Set WshShell = Wscript.CreateObject("Wscript.Shell")

if WshShell.AppActivate("Doki Doki Literature Club Plus") then
  WshShell.SendKeys("{UP}")
  WScript.Sleep(500)

  Dim LoopCounter
  LoopCounter = 0
  Do While LoopCounter < 200
    '
    ' DDLCを終了する
    '
    WshShell.SendKeys("{UP}")
    WScript.Sleep(500)
    WshShell.SendKeys("{Enter}")
    WScript.Sleep(500)
    WshShell.SendKeys("{Enter}")
    WScript.Sleep(4000)

    '
    ' DDLCを起動する
    '
    WshShell.SendKeys("{UP}")
    WScript.Sleep(1000)
    WshShell.SendKeys("{Enter}")
    WScript.Sleep(10000)

    LoopCounter = LoopCounter + 1
  Loop
end if
```

手順は以下のとおり。

* ドキドキ文芸部プラスを起動して、2 周目まですすめる
* 2 周目のメインメニューを表示する
* ゲームを起動したまま PowerShell 端末を起動する
* 前述のスクリプトをダウンロードする
* カレントディレクトリにスクリプトを配置する
* PowerShell 上で以下のコマンドを実行する

```ps1
wscript.exe dokidoki_literature_club.vbs
```

* 以降はひたすらゲーム上で勝手に操作されるので、終わるまで放置する。
  スクリプト実行中はその PC では他の操作をしてはいけないので、寝る直前に実行して放置するのがオススメ。
  * 解除できてなかったら、もう 1 回スクリプトを実行して放置する。
    そのうち解除されるはず。
* 以上

## 解説

シークレット 10 は 2 周目のメインメニューを表示したときに確率で取得できる。
よって、メインメニューをひたすら表示 → DDLC の終了を繰り返せばよい。
繰り返す操作はキーボード操作だけで完結できるので、繰り返すためのキーボード操作を VBScript で自動化した。

やってることは以下のとおり。

* ゲームのメインメニューを表示した直後、最初の 1 回だけ「カーソル上」を入力する
* 以降はループ処理
  * 次に「カーソル上」を入力する
  * Enter を押す
  * 「ゲームを終了するか？」的なポップアップが表示されるので Enter を入力して終了する
  * ドキドキ文芸部上のデスクトップが表示される
  * 「カーソル上」を押す
  * Enter を押す
  * メインメニューが表示される
* キーボード入力とキーボード入力の間には待ち時間が必要なので、適切な待ち時間を随時挟むようにしてる

200 回という数値は適当で、適当に 200 回にして動かしてる最中にシークレットを取得できたのでそのままにしている。

## 余談

何年前か忘れたけれど Steam で [Ampu-Tea](https://store.steampowered.com/app/289090/AmpuTea/?l=japanese) というゲームを買った。

このゲームは実績全解除が簡単だったので解除したけれど、1 個だけめちゃくちゃ面倒な実績があった。
内容は「500 回ゲームを起動する」というもの。

Ampu-Tea は正直に言うとクソゲーなので、500 回も遊ぶ価値は無い。
500 回も起動するのはさすがに面倒だった。
これを突破するために VBScript を書いた経験があって、それが今回も役に立った。

当該スクリプトがリポジトリに残っていたので、これを参考に今回のスクリプトを作成した。

* [start_amptea.vbs](https://github.com/jiro4989/sandbox/blob/a999bceb1979c5178747ff24a13143a91ecbd0a3/vbs/start_amptea.vbs)

このスクリプト単体だと、1 回ゲームを起動して終了しかしてない気がするので、この処理だけだと 500 回の実績を達成できないはず。
たぶん PowerShell 側でループ実行して 500 回繰り返したんだと思う。
