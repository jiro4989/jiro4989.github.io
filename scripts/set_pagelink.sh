#!/bin/bash

set -eu

find _posts -name '*.md' | generate_markdown_links | embed_links
