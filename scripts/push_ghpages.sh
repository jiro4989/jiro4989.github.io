#!/bin/bash

set -eu

git push
git switch gh-pages
git merge master
git push
git switch master
