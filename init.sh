#!/bin/bash
echo "现在拉取分支代码"
git fetch -p
echo "拉取代码"
git pull

echo "更新模块版"
go mod tidy
go mod download

