#!/bin/bash
# 餐厅点餐系统启动脚本（macOS/Linux）
echo "正在启动餐厅点餐系统..."
echo ""
# 检查是否有可执行权限，如果没有则添加
if [ ! -x "$(pwd)/order5003" ]; then
  chmod +x "$(pwd)/order5003"
  echo "已添加可执行权限"
fi
# 启动应用程序
"$(pwd)/order5003"
# 如果程序意外退出，显示错误信息
echo ""
echo "系统启动失败，请确保您有正确的可执行文件。"
echo "按Enter键退出..."
read