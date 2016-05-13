#!/usr/bin/env bash
curDir=$(cd `dirname $0`; pwd)

if [ ! -d "$curDir/17mei" ]; then
    git clone git@github.com:wothing/17mei.git
else
    pushd $curDir/17mei        
    git pull
fi

docker run -it --rm --net=ci -v /var/run/docker.sock:/var/run/docker.sock -v /root/.ssh/:/root/.ssh/ -v $curDir/woci.json:/woci.json -v $curDir/17mei:/gopath/src/github.com/wothing/17mei  index.tenxcloud.com/izgnod/dockerci
