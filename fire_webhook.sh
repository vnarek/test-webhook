#!/bin/bash

git reset HEAD~
git add -A
git commit -m "firing webhook"
git push -f origin master
echo "echoe";
