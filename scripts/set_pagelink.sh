#!/bin/bash

set -eu

if ! command -v generate_markdown_links; then
  (cd scripts/generate_markdown_links && go install)
fi

if ! command -v embed_links; then
  (cd scripts/embed_links && go install)
fi

find _posts -name '*.md' | generate_markdown_links | embed_links
