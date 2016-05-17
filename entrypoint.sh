#!/bin/bash
appDir=/app
ciDir=/ci
logDir=/log

###### start buildlog #####

pushd /
$appDir/buildlog.exe &


##### start webhook ######
pushd /
$appDir/webhook
