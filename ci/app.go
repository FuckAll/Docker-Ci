/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/13 22:10
 */

package ci

import (
	"strings"
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

	appDocker()
}

func builder(wg *sync.WaitGroup, jobs <-chan string) {
	for j := range jobs {
		CMD(j)
		wg.Done()
	}
}

func appDocker() {
	for _, s := range conf.Services {
		v := FMT("docker run -it -d --net=test -v /app:/app --name %s-%s alpine /app/%s %s", conf.Tracer, s.Name, s.Name, s.Para)
		v = strings.Replace(v, "[TRACER]", conf.Tracer, -1)
		CMD(v)
	}
}
