#!/usr/bin/env bash

# 使用方法：
# ./gen_model.sh -db=recharge_db -table=user,user_auth
# 再将./genModel下的文件剪切到对应服务的model目录里面，记得改package

#包名
modelPkgName=model
#表生成的genmodel目录
outPath="./"
# 数据库配置
host="127.0.0.1"
port=3306
dbname=""
username=root
passwd=123456

# 解析命令行参数
while [ $# -gt 0 ]; do
  case "$1" in
    -db=*)
      dbname="${1#*=}"
      ;;
    -table=*)
      tables="${1#*=}"
      ;;
    *)
      echo "Invalid argument: $1"
      exit 1
  esac
  shift
done

if [ -z "$dbname" ] || [ -z "$tables" ]; then
  echo "Usage: $0 -db=<dbname> -table=<table1,table2>"
  exit 1
fi

IFS=',' read -ra tablesArray <<< "$tables"
for table in "${tablesArray[@]}"; do
  echo "开始创建库：$dbname 的表：$table"
  gentool -dsn "${username}:${passwd}@tcp(${host}:${port})/${dbname}?charset=utf8mb4&parseTime=True&loc=Local" -tables "${table}" -onlyModel -modelPkgName="${modelPkgName}" -outPath="${outPath}"
done
