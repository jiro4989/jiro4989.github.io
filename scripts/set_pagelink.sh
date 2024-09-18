#!/bin/bash

set -eu

find _posts -name '*.md' |
  go run scripts/generate_markdown_links/main.go |
  go run scripts/embed_links/main.go
