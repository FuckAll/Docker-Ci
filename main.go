/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/13 16:50
 */

package main

import (
	"sync"

	"github.com/FuckAll/Docker-Ci/ci"
	_ "github.com/FuckAll/Docker-Ci/conf"
)

type StageFunc func() error

var wg = &sync.WaitGroup{}
var gofunc = func(foo StageFunc) {
	defer wg.Done()
	foo()
}

func main() {
	wg.Add(3)
	go gofunc(ci.Redis)
	go gofunc(ci.Pgsql)
	go gofunc(ci.Consul)
	wg.Wait()
	var err error
	if err = ci.AppBuild(); err != nil {
		ci.Clean()
	}
	if err = ci.AppTest(); err != nil {
		ci.Clean()
	}
	ci.Clean()
}
