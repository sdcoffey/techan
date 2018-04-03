#!/bin/bash

set -euf -o pipefail

make test

golint_count=`golint | wc -l | tr -d '[:space:]' || true`

if [[ $golint_count -gt 0 ]]; then
  echo "golint failed. Please fix the following issues:"
  golint
  exit 1
fi

echo -n "Version: "
read newversion

tag_count=`git tag --list | grep "$newversion" | wc -l | tr -d '[:space:]' || true`

if [[ $tag_count -gt 0 ]]; then
  echo "Tag: $newversion already exists"
  exit 1
else
  echo "Releasing $newversion"
fi

echo "Update CHANGELOG.md and press enter"
read

git add CHANGELOG.md

added_count=`git status --porcelain | grep "CHANGELOG.md" | wc -l | tr -d '[:space:]' || true`
if [[ $added_count -gt 0 ]]; then
  git commit -m"Release version $newversion"
fi

git tag "$newversion" -m "$newversion"

git push origin master
git push --tags
