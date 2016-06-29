/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/13 16:50
 */

package main

import (
	"flag"
	"fmt"
	"github.com/FuckAll/Docker-Ci/ci"
	"github.com/FuckAll/Docker-Ci/conf"
	"github.com/wothing/log"
	"strings"
)

var (
	onlyBuild   = flag.String("onlybuild", "", "only build app and image")
	testClean   = flag.Bool("testclean", false, "build test clean container")
	testNoClean = flag.Bool("testnoclean", false, "build test no clean container")
	push        = flag.Bool("push", false, "push image to repo")
	traceID     = flag.String("tid", "", "traceId for push")
	tag         = flag.String("tag", "", "image tag for push")
)

func main() {
	pre := `Docker-Ci Is A Tool Used to Build Docker Image And Push To Registry 
               Example:

                        1. Test All And Keep Docker Container
                          ./Docker-Ci -testnoclean 

                        2. Test All And Clean Docker Container
                          ./Docker-Ci -testclean  

                        3. ReBuild All Images
                          ./Docker-Ci -onlybuild [-tid 663d2166]

                        4. Build Some Images
                          ./Docker-Ci -onlybuild appway interway -tid 663d2166
               `
	fmt.Println(pre)
	flag.Parse()
	if *testClean {
		ci.CiRun("TestClean", "")
		return
	}
	if *testNoClean {
		ci.CiRun("TestNoClean", "")
		return
	}
	if *onlyBuild != "" {
		//如果单独的build某些微服务，是需要指定tid,否则会造成版本的错乱导致，无法测试
		app := strings.Split(*onlyBuild, ",")
		if len(app) > 0 {
			if *traceID != "" {
				conf.Tracer = *traceID
			} else {
				log.Tfatal(conf.Tracer, "TraceId is Empty!")
			}
			ci.CiRun("OnlyBuild", app...)
		} else {
			if *traceID != "" {
				conf.Tracer = *traceID
			}
			ci.CiRun("OnlyBuild")
		}

		return
	}
	if *push {
		if *traceID == "" {
			log.Fatal("TraceId is Empty!")
		}
		if *tag == "" {
			log.Fatal("Tag is Empty!")
		}
		conf.Tracer = *traceID
		ci.CiRun("Push", *tag)
		return
	}
}
