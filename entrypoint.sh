#!/bin/bash
appDir=/app
ciDir=/ci
logDir=/log

###### start buildlog #####

pushd /
/buildlog.exe &


##### start webhook ######
pushd /
$appDir/webhook
