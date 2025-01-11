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

<!-- textlint-disable -->
{% assign before_year = "" %}
{% for post in site.posts %}{% assign year = post.date | date: "%Y" %}{% if year != before_year %}

### {{ year }} 年
{% endif %}{% assign before_year = post.date | date: "%Y" %}
* {{ post.date | date: "%Y-%m-%d" }} {{ post.categories[0] }} [{{ post.title }}]({{ post.url }}){% endfor %}
<!-- textlint-enable -->
