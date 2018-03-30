#!/bin/bash

set -euf -o pipefail

make test

echo -n "Version: "
read newversion

existing_tag_count=`g tag --list | grep 0.2.0 | wc -l | tr -d '[:space:]'`

if [[ $existing_tag_count -gt "0" ]]; then
  echo "Tag: $newversion already exists"
fi

echo "Update CHANGELOG.md" and press enter
read

git add -A
git commit -m"Release version $newversion"

git tag -m $newversion

git push origin master
git push --tags
