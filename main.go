/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/13 16:50
 */

package main

import (
	"flag"
	"github.com/FuckAll/Docker-Ci/ci"
	"github.com/FuckAll/Docker-Ci/conf"
	"github.com/wothing/log"
	"strings"
)

var (
	onlyBuild = flag.String("onlybuild", "", `
	 1. Build All Images
            ./Docker-Ci -onlybuild all [-tid 663d2166]
	 2. Buidl Some Images   
	    ./Docker-Ci -onlybuild appway,interway -tid 663d2166 
	 `)
	testClean = flag.Bool("testclean", false, `
	 1. Test All And Clean Docker Container
	    ./Docker-Ci -testclean
	 `)
	testNoClean = flag.Bool("testnoclean", false, `
	 1.  Test All And Keep Docker Container
	    ./Docker-Ci -testnoclean
	 `)
	push = flag.String("push", "", `
	 1.  Push All Images To Registry
	    ./Docker-Ci -push all -tag v1.2.1 [-tid 663d2166]
	 2. Push Some Images To Registry
	    ./Docker-Ci -push appway,interway -tag v1.2.1 -tid 663d2166 
	 `)
	traceID = flag.String("tid", "", "TraceId For Push")
	tag     = flag.String("tag", "", "Image Tag For Push")
)

func main() {
	flag.Parse()
	if *testClean {
		ci.CiRun("TestClean", "")
		return
	}
	if *testNoClean {
		ci.CiRun("TestNoClean", "")
		return
	}
	if *push != "" {
		if *traceID == "" {
			log.Fatal("TraceId is Empty!")
		}
		if *tag == "" {
			log.Fatal("Tag is Empty!")
		}
		conf.Tracer = *traceID
		if *push != "all" {
			app := strings.Split(*push, ",")
			ci.CiRun("Push", *tag, app...)
		} else {
			ci.CiRun("Push", *tag)
		}
		return
	}

	if *onlyBuild != "" {
		//如果单独的build某些微服务，是需要指定tid,否则会造成版本的错乱导致，无法测试
		if *onlyBuild != "all" {
			app := strings.Split(*onlyBuild, ",")
			if *traceID != "" {
				conf.Tracer = *traceID
			} else {
				log.Tfatal(conf.Tracer, "TraceId is Empty!")
			}
			ci.CiRun("OnlyBuild", "", app...)
		} else {
			if *traceID != "" {
				conf.Tracer = *traceID
			}
			ci.CiRun("OnlyBuild", "")
		}

		return
	}
}
