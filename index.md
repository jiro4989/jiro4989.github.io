# 次郎の貝塚

* Table of contents
{:toc}

## はじめに

次郎が管理しているブログです。

このサイトの運営方針は [About](/about) に記載しています。
次郎（管理人）のことは [Profile](/profile) に記載しています。

このサイトのソースコードは <https://github.com/jiro4989/jiro4989.github.io> で公開しています。

## 投稿記事

記事の内容での検索は[検索フォーム](https://github.com/search?q=repo%3Ajiro4989%2Fjiro4989.github.io+path%3A%2F%5E_posts%5C%2F%2F+&type=code)からどうぞ。

投稿記事のカテゴリーはとおりです。

1. 技術: 主にソフトウェアエンジニアリング関係。たまにそれ以外も含む
1. ゲーム: ゲーム関係。主に遊んだゲームの話
1. 映画: 見た映画の話
1. イラスト: 絵に関連する話。描いた絵の話が主
1. 雑記: 上記のいずれでもない話

{% assign before_year = "" %}
{% for post in site.posts %}
{% assign year = post.date | date: "%Y" %}
{% if year != before_year %}
### {{ year }} 年

{% endif %}
* {{ post.date | date: "%Y-%m-%d" }} {{ post.categories[0] }} [{{ post.title }}]({{ post.url }})
{% endfor %}

一旦句切れ目。

<!-- START_POSTS -->
### 2025 年

* 2025-01-09 技術 [Google Search Console にサイトマップが登録されないけれど気にせず放置している](/tech/2025/01/09/search-console-sitemap.html)
* 2025-01-01 技術 [Web アクセシビリティの視点で自分の Web アプリを見つめ直す](/tech/2025/01/01/wcag-my-app.html)

### 2024 年

* 2024-12-28 雑記 [2024年振り返り](/daily/2024/12/28/furikaeri-2024.html)
* 2024-11-30 ゲーム [ドラゴンクエスト3リメイククリアした](/game/2024/11/30/dq3-remake.html)
* 2024-09-29 技術 [本ブログを支える技術](/tech/2024/09/29/my-blog-tech.html)
* 2024-09-21 ゲーム [Enotria The Last Song をクリアした感想](/game/2024/09/21/enotria-the-last-song.html)
* 2024-07-30 ゲーム [モンスターハンターライズ サンブレイクをクリアした](/game/2024/07/30/monster-hunter-rise.html)
* 2024-07-05 技術 [setup-nim-action を v2 にバージョンアップした](/tech/2024/07/05/setup-nim-action-v2.html)
* 2024-04-21 技術 [大量の IP が特定の CIDR に含まれているか調べるツールを作った](/tech/2024/04/21/check-ip-cidr.html)
* 2024-01-12 技術 [winget でパッケージを管理する](/tech/2024/01/12/manage-package-with-winget.html)
* 2024-01-06 雑記 [Twitter やめたけれど特に困っていない](/daily/2024/01/06/no-twitter.html)

### 2023 年

* 2023-12-12 雑記 [2023 年振り返り](/daily/2023/12/12/furikaeri-2023.html)
* 2023-12-06 技術 [Node プロジェクトで、ビルドプロセスでのみ必要なパッケージを dependencies で管理するのは誤り](/tech/2023/12/06/node-dependencies-ci.html)
* 2023-12-03 技術 [Bloodborne のビルドシミュレータを Next.js で作った](/tech/2023/12/03/bloodborne-build-simulator-next-js.html)
* 2023-11-19 技術 [メモを書く環境](/tech/2023/11/19/memo-environment.html)
* 2023-11-05 技術 [Rubyワンライナーの基本とユースケース](/tech/2023/11/05/ruby-oneliner.html)
* 2023-10-28 ゲーム [エルデンリングの自作Lv150ビルドと感想](/game/2023/10/28/eldenring-build.html)
* 2023-10-15 映画 [ホラー映画の感想](/movie/2023/10/15/movie.html)
* 2023-10-15 雑記 [ブログ移転しました](/daily/2023/10/15/blog-changelog.html)

### 2022 年

* 2022-12-16 雑記 [2022 年振り返り](/daily/2022/12/16/furikaeri-2022.html)
* 2022-10-20 雑記 [Kinesis Advantage 360 を買った](/daily/2022/10/20/kinesis-advantage-360.html)
* 2022-08-06 技術 [AWS ソリューションアーキテクトアソシエイト試験に合格した](/tech/2022/08/06/aws-saa.html)
* 2022-07-19 イラスト [壱百満天原サロメお嬢様描いた](/illust/2022/07/19/illust-100mantenbara-salome.html)
* 2022-06-28 イラスト [グリザイユ画法で絵を描いてみた](/illust/2022/06/28/illust-grisaille.html)
* 2022-06-19 技術 [テキストを壱百満天原サロメお嬢様文体に変換するコマンドを書いた](/tech/2022/06/19/ojosama.html)
* 2022-02-05 技術 [ドキドキ文芸部プラス！のシークレット 10 を取得するための VBScript を書いた](/tech/2022/02/05/dokidoki-literature-club-vbscript.html)
* 2022-02-01 雑記 [親知らずの抜歯と抜歯後の経過](/daily/2022/02/01/oyashirazu.html)

### 2021 年

* 2021-12-14 技術 [色彩検定 2 級に合格した](/tech/2021/12/14/shikisai-kentei-2-kyu.html)
* 2021-12-07 雑記 [2021 年振り返り](/daily/2021/12/07/furikaeri-2021.html)

### 2020 年

* 2020-12-31 雑記 [2020 年振り返り](/daily/2020/12/31/furikaeri-2020.html)
* 2020-07-14 技術 [rpm パッケージを作成する GitHub Actions を作った](/tech/2020/07/14/github-actions-rpm-package.html)
* 2020-07-10 技術 [debian パッケージを作成する GitHub Actions を作った](/tech/2020/07/10/github-actions-debian-package.html)

### 2019 年

* 2019-12-29 雑記 [2019 年振り返り](/daily/2019/12/29/furikaeri-2019.html)
* 2019-04-01 雑記 [新卒入社してから 2 年勤めた SIer を退職した](/daily/2019/04/01/taishoku-entry-shinsotsu.html)

### 2018 年

* 2018-12-29 雑記 [2018 年振り返り](/daily/2018/12/29/furikaeri-2018.html)

<!-- END_POSTS -->
