#!/bin/bash
# 餐厅点餐系统编译脚本
# 此脚本将为不同操作系统编译可执行文件

echo "开始编译餐厅点餐系统..."

echo "\n编译macOS版本..."
go build -o order5003-macos cmd/api/main.go
if [ $? -eq 0 ]; then
echo "macOS版本编译成功！"
else
echo "macOS版本编译失败！"
fi

echo "\n编译Linux版本..."
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o order5003-linux cmd/api/main.go
if [ $? -eq 0 ]; then
echo "Linux版本编译成功！"
else
echo "Linux版本编译失败！"
fi

echo "\n编译Windows版本..."
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o order5003.exe cmd/api/main.go
if [ $? -eq 0 ]; then
echo "Windows版本编译成功！"
else
echo "Windows版本编译失败！"
fi

echo "\n编译完成！生成的文件："
ls -la order5003* | grep -v "\.sh\|\.bat"

echo "\n提示：请将对应系统的可执行文件和启动脚本一起提供给用户使用。"

echo "\n按Enter键退出..."
read