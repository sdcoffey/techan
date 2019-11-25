#!/bin/bash

set -euf -o pipefail

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

git tag "v$newversion" -m "v$newversion"

git push origin master
git push --tags
