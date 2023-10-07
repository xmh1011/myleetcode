#!/bin/bash

file_path="/Users/xiaominghao/code/myleetcode"

readme_path="$file_path/README.md"

log_path="$file_path/log.txt"

now=$(date +"%Y-%m-%d %H:%M:%S")

date=$(date +"%Y-%m-%d")

line_number=$(($(wc -l < "$readme_path") - 4))

title="- 🗓 Automatically updated by [script](./script.sh) at: "

new_content="$title$now"

sed -i '' "${line_number}s/.*/$new_content/" "$readme_path"

export https_proxy=http://127.0.0.1:7890 http_proxy=http://127.0.0.1:7890 all_proxy=socks5://127.0.0.1:7890

cd $file_path

echo "success: $now" >> "$log_path"

git add .
git commit -s -m "feat: $date commit"
git push origin main -f