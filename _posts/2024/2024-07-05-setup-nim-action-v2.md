---
layout: default
title: "setup-nim-action を v2 にバージョンアップした"
date: 2024-07-05 09:00:00 +0900
categories: tech
categoriesJP: 技術
---

# setup-nim-action を v2 にバージョンアップした

[setup-nim-action](https://github.com/jiro4989/setup-nim-action) という自作の GitHub Actions（カスタムアクション）がある。

2019年ごろに公開してからほそぼそと更新を続けていて、おそらくGitHubのNimプロジェクトのCIで一番使われている Nim setup アクションなんじゃないかと自負してる。
もっとも、ここ数年は特に大きなアップデートもなく、node\_modules の依存ライブラリ更新しかしてなかったので、大きな変更がない状態だった。

そんな setup-nim-action がとある理由で大改修することになり、生まれ変わってメジャーバージョンアップした。

* Table of contents
{:toc}

## 突然 Nim のインストールが遅くなった

ある日いきなり CI が急激に遅くなった。
今までは10秒もかかっていなかったコンパイラのインストールに6分くらいかかるようになった。
Macに至っては30分もかかるようになり、issues が作られた。

* <https://github.com/jiro4989/setup-nim-action/issues/483>

特に僕はsetup-nim-actionに手を入れていなかったのにいきなり CI が遅くなったので、原因があるとすれば以下のどちらかだろうと考えた。

* GitHub Actions のランナーに変更が入った
* 内部で使っている [choosenim](https://github.com/dom96/choosenim) に変更が入った

しかし choosenim も新バージョンは特に公開されていないので、挙動が変わるようには思えなかった。

正直、何が原因なのか皆目検討がつかなかったのでお手上げ状態だったのだけれど、とはいえユーザが多いので放置するわけにもいかなかった。
そんなわけでとにかくインストールを高速で終わらせるためにあらゆる調査をした。

## 調査と修正

結果としては、choosenim を使ってインストールしようとすると非常に時間がかかるとわかった。
それもGitHub Actions上で実行した場合のみ。

choosenim を使うと node ランタイムでも composite アクションでもインストールに時間がかかっていた。
setup-nim-action v1 はTypeScriptで実装しており、node20 ランタイムで動作していた。
その内部では choosenim に依存しており、アクション実行時に choosenim をダウンロードしてきて実行して Nim をインストールするような作りになっている。

GitHub Actionsのカスタムランタイムは、node 系と composite の 2 つが選択できる。
composite はあとから追加されたランタイムで、簡単にいうとGitHub Actionsのワークフローをかくときの YAML 構文がそのままカスタムアクションになる。
したがって、実装は bash あるいは PowerShellが基本になる。

今回の問題の調査の結果 **node ランタイムから composite に切り替えて、choosenim で実装されていた Nim をインストールする処理をすべてbashで実装しなおす** ことにした。

もちろん、choosenim のすべてをシェルスクリプトに移植したわけではない。
choosenim は実行プラットフォームを判定して必要な依存ライブラリを整えるような処理も含まれる。
setup-nim-action はそこまでやらず、Nimコンパイラのインストール部分のみを実装することにした。
依存ライブラリ（GCCとか）はすべて GitHub Actions のランナー側でほぼ揃うだろう、と考えたため。

この場合、curl でビルド済みのコンパイラを落としてきて配置してPATHを通すだけで良い。
bashで実装してもそこまで膨大なコードにはならないだろうと判断した。

対応したPRは以下。

* <https://github.com/jiro4989/setup-nim-action/pull/491>

choosenim にはなくて setup-nim-action で独自に実装していた `1.x` とか `1.6.x` とかで最新版を取ってくる処理もシェルスクリプトで実装しなおした。
この対応により、setup-nim-action は約5年間ずっと依存していた choosenim への依存がなくなった。

そしてNimのインストール速度はいぜんと同等の速度でインストールできる状態に戻った。

TypeScriptなども不要になったため、リポジトリ上から node.js, typescript のコードがすべて消え、node\_modules も不要になった。
node ランタイムのカスタムアクションでは、node_modules をリポジトリのコードとして抱えなければならない制約があるので、コード量がめちゃくちゃ多くなる。

結果として、setup-nim-action はたかだか130行程度のシェルスクリプトと、少しの action.yml だけで動作する非常にコンパクトなカスタムアクションへと生まれ変わった。

## 破壊的変更とマイグレーション

僕はchoosenimの中身に詳しいわけではないため、choosenimの移植をしたわけではない。

やったこととしては、GitHub Actions上でNimが使えるように必要最低限の処理のみをシェルスクリプトで実装した。
かなり突貫工事での実装になったし、ランタイムも変更になった。Nimのインストール先も変更せざるを得なかった。

ActionsのInputの必須パラメータは変えないように実装したので、使い勝手は今までと変わらないものの、大規模な変更になってしまった。
そのため、苦渋の決断でメジャーバージョンアップして v2 バージョンを発行した。

破壊的変更をするならマイグレーションガイドが必要だろうと思い、かなり詳細にマイグレーションガイドを書いた。
ついでだったので、[Nim wikiにかかれてるCI周り](https://github.com/nim-lang/Nim/wiki/BuildServices)も更新した。

## やり終えた感想

急な対応だったのでかなり心身ともに疲れた。

が、結果としてリポジトリは非常にコンパクトになって、依存パッケージ管理からも開放されて非常に気分が良い。
dependabot で node\_modules 更新するだけの運用から開放された。
choosenim への依存をなくしたかったのも実現できた。
昔からどうにかしたいと思っていたので。

何気に composite アクションを初めて使ったので新鮮だった。
シェルスクリプトだけの簡単なカスタムアクション作るときは、全部これで良いと思う。

## setup-nim-actionはnim-langグループの管理下になってほしい

なんだかんだ setup-nim-action の保守を続けることもう5年。
事実上 Nim プロジェクトで一番使われてる GitHub Actions（のはず）なんだし GitHub の nim-lang グループの管理下になってほしい気持ちがある。
僕一人で保守しつづけてもまぁいいんだけれど、僕が突然亡くなったりしたら保守できる人いなくなっちゃうし。
最近は Nim のコードほとんど書いてないのもあって、Nim公式が保守してほしい気持ちがある。

ただプロジェクトの移管は、問題もある。
つい最近 Polyfill が売却されてマルウェアを仕込まれた事件があるので、どこまで信じていいのか分からない。

* [マルウェア混入が発覚したJavaScriptライブラリ「Polyfill.io」のドメインを登録事業者が停止 - Gigazine](https://gigazine.net/news/20240702-namecheap-polyfill-io-supply-chain-attack/)

まぁさすがに Nim 公式がそういうことするとは思えないけどね。
