#!/bin/bash

set -eu

read -r -p "タイトルを入力してください: " title
read -r -p "ファイル名を入力してください: " filename
read -r -p "日付を入力してください(yyyy-mm-dd): " dt
read -r -p "カテゴリーを入力してください(game, tech, daily, movie): " category

post_file="_posts/${dt}-${filename}.md"

cat << EOS > "${post_file}"
---
layout: default
title: "${title}"
date: ${dt} 09:00:00 +0900
categories: ${category}
---

# ${title}

TODO

* Table of contents
{:toc}

## TODO

TODO
EOS

echo "${post_file} が作成されました"
