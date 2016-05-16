#!/usr/bin/env bash
curDir=$(cd `dirname $0`; pwd)

if [ ! -d "$curDir/17mei" ]; then
    git clone git@github.com:wothing/17mei.git
else
    pushd $curDir/17mei        
    git pull
fi

num=`docker network ls | awk '{print $2}' | grep test | wc -l`
[ $num -ge 1 ] || docker network create test

docker run -it --rm --net=test -v /var/run/docker.sock:/var/run/docker.sock -v /root/.ssh/:/root/.ssh/ -v $curDir/17mei:/gopath/src/github.com/wothing/17mei -v /root/.bashrc:/root/.bashrc -v app:/app -v log:/log/ -v $curDir/hooks.json:/hooks.json -v $curDir/webhook.sh:/webhook.sh -v $curDir/woci.json:/woci.json -v $curDir/buildlog.exe:/buildlog.exe  -p 9090:9090 -p 9000:9000 index.tenxcloud.com/izgnod/dockerci
