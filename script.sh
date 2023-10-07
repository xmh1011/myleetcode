#!/bin/bash

file_path="/Users/xiaominghao/code/myleetcode/README.md"

now=$(date +"%Y-%m-%d %H:%M:%S")

date=$(date +"%Y-%m-%d")

line_number=$(($(wc -l < "$file_path") - 4))

title="- 🗓 Updated at: "

new_content="$title$now"

sed -i '' "${line_number}s/.*/$new_content/" "$file_path"

export https_proxy=http://127.0.0.1:7890 http_proxy=http://127.0.0.1:7890 all_proxy=socks5://127.0.0.1:7890

cd /Users/xiaominghao/code/myleetcode

git add .
git commit -s -m "feat: $date commit"
git push origin main -f

echo "success: $now" >> log.txt