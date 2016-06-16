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
	"github.com/wothing/log"
)

var (
	onlyBuild   = flag.Bool("onlybuild", false, "only build app and image")
	testClean   = flag.Bool("testclean", false, "build test clean container")
	testNoClean = flag.Bool("testnoclean", false, "build test no clean container")
	push        = flag.Bool("push", false, "push image to repo")
	traceID     = flag.String("tid", "", "traceId for push")
	tag         = flag.String("tag", "", "image tag for push")
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
	if *onlyBuild {
		ci.CiRun("OnlyBuild", "")
		return
	}
	if *push {
		if *traceId == "" {
			log.Fatal("TraceId is Empty!")
		}
		if *tag == "" {
			log.Fatal("Tag is Empty!")
		}
		ci.CiRun("Push", *traceId, *tag)
		return
	}
}
