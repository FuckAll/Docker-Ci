/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/13 09:46
 */

package ci

import (
	. "github.com/FuckAll/Docker-Ci/cmdrun"
	"github.com/FuckAll/Docker-Ci/conf"
)

func AppTest() {
	CMD(FMT("CGO_ENABLED=0 go test -c -o /app/testbin %s/gateway/tests/*.go", conf.ProjectPath))
	CMD(FMT("TestEnv=CI CiTracer=%s /app/testbin -test.v ", conf.Tracer))
}
