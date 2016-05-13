#!/usr/bin/env bash
curDir=$(cd `dirname $0`; pwd)

if [ ! -d "$curDir/17mei" ]; then
    git clone git@github.com:wothing/17mei.git
else
    pushd $curDir/17mei        
    git pull
fi

docker run -it --rm --net=ci -v /var/run/docker.sock:/var/run/docker.sock -v app:/app -v $curDir/:/Ci  dockerci
