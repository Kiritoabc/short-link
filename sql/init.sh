#!/bin/bash

# 创建数据库 请自行修改数据库密码等信息
mysql -uroot -p123456 -e "CREATE DATABASE IF NOT EXISTS shortlink DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;"

# 导入SQL文件 请自行修改SQL文件路径
mysql -uroot -p123456 shortlink < short_link.sql

# 提示完成
echo "Database import complete."