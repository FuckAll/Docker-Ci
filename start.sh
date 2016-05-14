#!/usr/bin/env bash
curDir=$(cd `dirname $0`; pwd)

if [ ! -d "$curDir/17mei" ]; then
    git clone git@github.com:wothing/17mei.git
else
    pushd $curDir/17mei        
    git pull
fi

docker network create test 

docker run -it --rm --net=test -v /var/run/docker.sock:/var/run/docker.sock -v /root/.ssh/:/root/.ssh/ -v $curDir/woci.json:/woci.json -v $curDir/17mei:/gopath/src/github.com/wothing/17mei  -v $curDir/bin/linux_64:/ci  -v /root/.bashrc:/root/.bashrc index.tenxcloud.com/izgnod/dockerci
