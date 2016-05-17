#!/bin/bash

WORK_DIR=/gopath/src/github.com/wothing/17mei
Ci_DIR=/ci/

[ $1 ] || exit 1
[ $2 ] || exit 1

Git_Url=$1
Full_Name=$2

//获取最新代码
pushd $WORK_DIR
git checkout develop
git pull 

//添加本地测试分支
git checkout -b $2-develop develop
git pull $1 develop


//Ci 全部的代码
pushd $Ci_DIR
/app/Docker-Ci 
