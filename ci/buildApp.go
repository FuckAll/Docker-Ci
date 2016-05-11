/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/05 14:15
 */

package ci

import (
	// "strings"
	"sync"

	"github.com/wothing/log"

	. "github.com/FuckAll/Docker-Ci/cmdrun"
	"github.com/FuckAll/Docker-Ci/conf"
)

func AppBuild() {
	CMD("make -C " + conf.ProjectPath + " idl ve")

	jobCount := len(conf.Services)
	jobs := make(chan string, jobCount)

	wg := &sync.WaitGroup{}
	wg.Add(jobCount)

	for i, j := 0, conf.Concurrent; i < j; i++ {
		go builder(wg, jobs)
	}

	//add jobs
	for _, s := range conf.Services {
		jobs <- FMT("cd %s/%s && CGO_ENABLED=0 go install", conf.ProjectPath, s.Path)
	}

	wg.Wait()
	log.Tinfo(conf.Tracer, "All build job done")
}

func builder(wg *sync.WaitGroup, jobs <-chan string) {
	for j := range jobs {
		CMD(j)
		wg.Done()
	}
}
