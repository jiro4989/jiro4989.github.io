#!/bin/bash

set -eu

declare -A table=(
  ["tech"]="技術"
  ["movie"]="映画"
  ["game"]="ゲーム"
  ["daily"]="雑記"
)

for f in _posts/*.md; do
  title="$(grep -m1 -E '^title:' "$f" | awk -F : '{print $2}' | tr -d '"' | sed -E "s/^\s+//")"
  url_suffix="$(basename "${f:18}" .md)"
  time="$(grep -m1 -E '^date:' "$f" | awk '{print $2}')"
  time2="$(tr - / <<< "$time")"
  sortkey="$(grep -m1 -E '^date:' "$f" | awk '{print $2, $3, $4}')"
  cat="$(grep -m1 -E '^categories:' "$f" | awk '{print $2}')"
  cat2="${table["$cat"]}"
  echo "$sortkey * $time2 $cat2 [$title](/$cat/$time2/${url_suffix}.html)"
done |
  sort -r |
  sed -E 's/^[^\*]+//' |
  ./scripts/embed_links.rb
