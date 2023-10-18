#!/bin/bash

set -eu

info() {
  echo "[$(date +'%Y-%m-%d %H:%M:%S')] INFO | $*"
}

if [[ $# -eq 0 ]] || [[ $1 =~ ^(-h|--help)$ ]]; then
  cat << EOS
setdate.sh は _posts 配下の markdown の date パラメータを更新して、ファイル名の日付を更新するスクリプトである。

usage:
    setdate.sh <_posts/filepath.md>
    setdate.sh <-h | --help>

options:
    -h, --help    このヘルプを出力する。
EOS
  exit
fi

target_file=$1
now="$(date +'%Y-%m-%d %H:%M:%S') +0900"
ymd="$(echo "$now" | awk '{print $1}')"
dest_file="$(echo "$target_file" | sed -E "s/[0-9]{4}-[0-9]{2}-[0-9]{2}/$ymd/")"

sed -Ei "/---/,/---/s/^date:.*/date: $now/" "$target_file"
info "updated date of '$target_file'"

mv "$target_file" "$dest_file"
info "moved '$target_file' to '$dest_file'"

info "script completed"
