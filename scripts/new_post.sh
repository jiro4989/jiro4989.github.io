#!/bin/bash

set -eu

err() {
  echo "$(date +"%Y-%m-%d %H:%M:%S") [ERR] $*" >&2
}

read -r -p "タイトルを入力してください: " title

read -r -p "ファイル名を入力してください: " filename
if ! [[ "$filename" =~ ^[-a-z0-9]+$ ]]; then
  err "ファイル名は a-z 0-9 - のみ使用可能です: filename = $filename"
  exit 1
fi

read -r -p "日付を入力してください(yyyy-mm-dd): " dt
if ! [[ "$dt" =~ ^2[0-9]{3}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$ ]]; then
  err "日付は yyyy-mm-dd です: dt = $dt"
  exit 1
fi

read -r -p "カテゴリーを入力してください(game, tech, daily, movie): " category
case $category in
  game | tech | daily | movie)
    # 正常系なのでなにもしない
    ;;
  *)
    err "カテゴリーは game, tech, daily, movie のみ指定可能です: category = $category"
    exit 1
    ;;
esac

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
