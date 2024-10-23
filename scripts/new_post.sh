#!/bin/bash

set -eu

err() {
  echo "$(date +"%Y-%m-%d %H:%M:%S") [ERR] $*" >&2
}

title=""
filename=""
dt=""
category=""
overview="TODO"

if [[ $# -eq 1 ]] && [[ $1 = furikaeri ]]; then
  read -r -p "何年の振り返りか入力してください(yyyy): " year
  title="${year} 年振り返り"
  filename="furikaeri-${year}"
  category="daily"
  overview="${year} 年も終わるので、今年 1 年を振り返る。"
fi

if [[ "$title" = "" ]]; then
  read -r -p "タイトルを入力してください: " title
fi

if [[ "$filename" = "" ]]; then
  read -r -p "ファイル名を入力してください: " filename
fi
if ! [[ "$filename" =~ ^[-a-z0-9]+$ ]]; then
  err "ファイル名は a-z 0-9 - のみ使用可能です: filename = $filename"
  exit 1
fi

if [[ "$dt" = "" ]]; then
  read -r -p "日付を入力してください(yyyy-mm-dd): " dt
fi
if ! [[ "$dt" =~ ^2[0-9]{3}-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])$ ]]; then
  err "日付は yyyy-mm-dd です: dt = $dt"
  exit 1
fi

if [[ "$category" = "" ]]; then
  read -r -p "カテゴリーを入力してください(game, tech, daily, movie, illust): " category
fi
case $category in
  game | tech | daily | movie | illust)
    # 正常系なのでなにもしない
    ;;
  *)
    err "カテゴリーは game, tech, daily, movie, illust のみ指定可能です: category = $category"
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

${overview}

* Table of contents
{:toc}

## TODO

TODO
EOS

echo "${post_file} が作成されました"
