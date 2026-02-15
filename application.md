# 自作アプリ

自作のアプリ、Web サイト、便利ツールについて。
数が多いので特筆したいものだけ書いている。
基本的にソースコードはすべて [GitHub](https://github.com/jiro4989) で管理している。

* Table of contents
{:toc}

## Web アプリ

### websh

<https://github.com/jiro4989/websh>

シェル芸 Bot の Web アプリ。内部ではシェル芸 Bot の Docker コンテナが起動してシェルを実行する。
2019 年〜 2024 年の間運用していたが、サーバの保守を続けるのが大変になって停止した。

### ojosama web converter

<https://github.com/jiro4989/ojosama-web>

テキストを入力すると、Vtuber の壱百満天原サロメお嬢様っぽい言い回しに変換して返却するジョークアプリ。
インフラは AWS CloudFront + API Gateway + Lambda で動かしている（非公開）。
最初は Heroku でアプリを稼働していたが、Heroku 無料枠の廃止に伴って AWS に移管して今の構成になった。

### jira\_issue\_url\_generator

<https://github.com/jiro4989/jira_issue_url_generator>

Jira の issue url を生成するためのツール。
仕事で Jira を使うことが多いので自分用に作った。
結構古いツールで、当時興味があった React.js の勉強も兼ねて実装したもの。

### bloodborne-build-simulator

<https://github.com/jiro4989/bloodborne-build-simulator>

Bloodborne というゲームのステータス計算機。
GitHub Pages でホストしている。
計算ロジックは簡単で実装に苦労しなさそうだったので、勉強もかねて以前から興味があった Next.js で実装した。

## CLI ツール

### textimg

<https://github.com/jiro4989/textimg>

ANSI エスケープシーケンスを解析して、それを色コードに変換して画像出力するツール。
ターミナルで色のついたテキストを、その見た目のまま画像出力する。
もともとはシェル芸 Bot 用に作り始めたものだけれど、いつの間にか汎用的なツールになった。

### ojosama

<https://github.com/jiro4989/ojosama>

上の方に書いた ojosama web converter のロジック部分＋CLI ツール。
壱百満天原サロメお嬢様が初登場した直後くらいに実装した。
最初はこのライブラリ＋CLI だけ実装していたが、その後 Web アプリ版も実装して公開した。

### slotchmod

<https://github.com/jiro4989/slotchmod>

スロットマシーンみたいにファイルパーミッションを設定するジョークツール。
[ファイルパーミッションでスロットがしたい - Qiita](https://qiita.com/jiro4989/items/2530c4f789916521a47a)って記事をアドベントカレンダー 1 日目に投稿してバズったりした。
こういうしょうもないツールを真面目に作ってるときが最高に楽しい。

## GUI ツール

### TKoolFacetileMaker2

<https://github.com/jiro4989/TKoolFacetileMaker2>

RPG ツクール用の画像タイル生成ツール。
Kotlin + JavaFX で作った GUI アプリ。
最近は GUI デスクトップアプリを作らなくなったので、自作アプリの中でも珍しい部類。

初期バージョンは Java swing で作っていて、その後 JavaFX に変更した。
最初は Java で書いていたのも途中で Kotlin に変更したりと、大きな変更を何度もしてる。
2016 年から開発が続いている僕の最古参ツール。
