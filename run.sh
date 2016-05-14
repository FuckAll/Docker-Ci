#!/bin/bash


## start webhook and buildlog ##
/app/buildlog &
/app/webhook -hooks="/ci/hooks.json"  
