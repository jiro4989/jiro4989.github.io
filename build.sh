#!/bin/bash

set -eu

build() {
  docker run --rm \
    -v $(pwd):/documents/ \
    asciidoctor/docker-asciidoctor \
    asciidoctor \
    -r asciidoctor-diagram \
    $1
}

# asciidocからhtmlを生成
for dir in page/blog/*; do
  (
    cd "$dir"
    for file in *.adoc; do
      build "$file"
    done
  )
done

# リリース用のディレクトリに移動
find page/blog/2019/*.html -type f -exec mv -f {} docs/blog/2019/ \;
sudo chown -R $USER. docs
