#!/usr/bin/env bash


## 获取变量个数
if [ $# -lt 1 ]; then
    echo "error.. need tid"
    exit 1
fi

## 停止正在使用删除镜像的服务
containers=$(docker ps -a |grep $1| awk '{print $1}')
for c in $containers; do
    docker stop $c
    docker rm $c
done

## 删除操作
images=$(docker images -a | awk '{image=$1":"$2; print image}' | grep $1)
for i in $images; do
    docker rmi $i
done