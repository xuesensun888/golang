#!/bin/bash

# 要监视的目录路径
WATCH_DIR="/opt/cargo"

# 监听目录中的文件创建事件
inotifywait -m -e create --format '%w%f' "$WATCH_DIR" | while read FILE
do
	    echo "New file created: $FILE"
	        # 在此处添加要执行的操作或脚本
	done

