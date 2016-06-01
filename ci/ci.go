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
	"github.com/FuckAll/Docker-Ci/test"
	"github.com/wothing/log"
)

func CiRun(step string, args ...string) {
	switch step {
	case "OnlyBuild":
		log.Tinfo(conf.Tracer, "OnlyBuild Start!")
		CiBuildApp()
		log.Tinfo(conf.Tracer, "OnlyBuild Complate!")
	case "TestClean":
		log.Tinfo(conf.Tracer, "TestClean Start!")
		CiTestAppClean()
		log.Tinfo(conf.Tracer, "TestClean Complate!")
	case "TestNoClean":
		log.Tinfo(conf.Tracer, "TestNoClean Start!")
		CiTestAppNoClean()
		log.Tinfo(conf.Tracer, "TestNoClean Complate!")
	case "Push":
		log.Tinfo(conf.Tracer, "Push Start!")
		if len(args) < 2 {
			log.Tfatal(conf.Tracer, "PushImage Cant't Get TraceId And Tag")
		}
		CiPush(args[0], args[1])
		log.Tinfo(conf.Tracer, "Push Complate!")
	default:
		fmt.Println("CiRun Do Nothing!!!")
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
	test.TestApp()

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
	test.TestApp()
	//5. Clean App
	err = container.RemoveAppContainer()
	if err != nil {
		log.Tfatalf(conf.Tracer, "Ci CleanConta Error: %s", err)
	}

}

func CiPush(traceId, tag string) {
	// 修改镜像Tag --> Push到Repo --> 删除旧镜像
	Registry := conf.Config.Registry
	for _, service := range conf.Config.Services {
		Name := traceId + "-" + service.Name
		Repo := Registry + "/" + Name
		Tag := tag
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
