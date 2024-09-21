#!/bin/bash

set -eu

(cd scripts/embed_links && go install)
(cd scripts/generate_markdown_links && go install)
