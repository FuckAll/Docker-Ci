/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/13 22:10
 */

package ci

import (
	"fmt"

	"github.com/FuckAll/Docker-Ci/api"
	"github.com/FuckAll/Docker-Ci/build"
	"github.com/FuckAll/Docker-Ci/conf"
	"github.com/FuckAll/Docker-Ci/container"
	"github.com/FuckAll/Docker-Ci/infrastructure"
	"github.com/wothing/log"
)

func CiRun(step, traceId string) {
	switch step {
	case "OnlyBuild":
		CiBuildApp()
	case "TestClean":
		CiTestAppClean()
	case "TestNoClean":
		CiTestAppNoClean()
		fmt.Println("Clean Complate")
	case "Push":
		if traceId == "" {
			log.Tfatal(conf.Tracer, "PushImage Cant't Get TraceId")
		}
		CiPush(traceId)
	default:
		fmt.Println("CiRun Complate!")
	}
}

func CiBuildApp() {
	_, err := build.BuildApp()
	if err != nil {
		log.Tfatalf(conf.Tracer, "BuildApp Error: %s", err)
	}
	build.CreateDockerFile()
	build.BuildImage()

}

func CiTestAppNoClean() {
	// 1. 构建镜像
	CiBuildApp()
	// 2. 启动基础服务，例如：pgsql redis consul
	err := infrastructure.StartConsul()
	if err != nil {
		log.Tfatalf(conf.Tracer, "Ci StartConsul Error: %s", err)
	}
	err = infrastructure.StartRedis()

	if err != nil {
		log.Tfatalf(conf.Tracer, "Ci StartRedis Error: %s", err)
	}
	err = infrastructure.StartPostgres()
	if err != nil {
		log.Tfatalf(conf.Tracer, "Ci StartPostgres Error: %s", err)
	}
	//3. 启动业务代码容器
	err = container.StartApp()
	if err != nil {
		log.Tfatalf(conf.Tracer, "Ci StartApp Error: %s", err)
	}
	//4. 测试

}

func CiTestAppClean() {
	CiBuildApp()
	//2. 启动基础服务，例如：pgsql redis consul
	err := infrastructure.StartConsul()
	if err != nil {
		log.Tfatalf(conf.Tracer, "Ci StartConsul Error: %s", err)
	}
	err = infrastructure.StartRedis()

	if err != nil {
		log.Tfatalf(conf.Tracer, "Ci StartRedis Error: %s", err)
	}
	err = infrastructure.StartPostgres()
	if err != nil {
		log.Tfatalf(conf.Tracer, "Ci StartPostgres Error: %s", err)
	}
	//3. 启动业务代码容器
	err = container.StartApp()
	if err != nil {
		log.Tfatalf(conf.Tracer, "Ci StartApp Error: %s", err)
	}
	//4. 测试
	//5. Clean

}

func CiPush(traceId string) {
	// 修改镜像Tag --> Push到Repo --> 删除旧镜像
	Registry := conf.Config.Registry
	for _, service := range conf.Config.Services {
		Name := traceId + "-" + service.Name
		Repo := Registry + "/" + Name
		Tag := traceId
		err := api.ChangeTag(Repo, Tag, Name)
		if err != nil {
			log.Tfatalf(conf.Tracer, "Ci CiPush ChangeTag  Error: %s", err)
		}
		err = api.PushImage(Repo, Tag, Registry)
		if err != nil {
			log.Tfatalf(conf.Tracer, "Ci CiPush PushImage  Error: %s", err)
		}
		err = api.RemoveImage(Name)
		if err != nil {
			log.Tfatalf(conf.Tracer, "Ci CiPush RemoveImage Error: %s", err)
		}
	}
}
