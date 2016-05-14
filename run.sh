#!/bin/bash


## start webhook and buildlog ##
mv /ci
/app/buildlog &
/app/webhook -hooks="/ci/hooks.json"  
