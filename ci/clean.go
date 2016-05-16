/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/13 07:27
 */

package ci

import (
	. "github.com/FuckAll/Docker-Ci/cmdrun"
	"github.com/FuckAll/Docker-Ci/conf"
)

func Clean() {
	CMD(FMT("docker stop %s-consul", conf.Tracer))
	CMD(FMT("docker rm %s-consul", conf.Tracer))

	CMD(FMT("docker stop %s-redis", conf.Tracer))
	CMD(FMT("docker rm %s-redis", conf.Tracer))

	CMD(FMT("docker stop %s-pgsql", conf.Tracer))
	CMD(FMT("docker rm %s-pgsql", conf.Tracer))

	for _, s := range conf.Services {
		CMD(FMT("docker stop %s-%s", conf.Tracer, s.Name))
		CMD(FMT("docker rm %s-%s", conf.Tracer, s.Name))
	}
}
