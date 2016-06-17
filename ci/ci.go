/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/13 22:10
 */

package ci

import (
	"fmt"
	"time"

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
	// 2. 启动基础服务，例如：pgsql redis etcd
	//	err := infrastructure.StartConsul()
	//if err != nil {
	//log.Tfatalf(conf.Tracer, "Ci StartConsul Error: %s", err)
	//	}
	err := infrastructure.StartEtcd()
	if err != nil {
		log.Tfatal(conf.Tracer, "Ci StartEtcd Error: %s", err)
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
	time.Sleep(10 * time.Second)
	err = container.StartApp()
	if err != nil {
		log.Tfatalf(conf.Tracer, "Ci StartApp Error: %s", err)
	}
	//4. 测试
	test.TestApp()

}

func CiTestAppClean() {
	// 1. 构建镜像
	CiBuildApp()
	//2. 启动基础服务，例如：pgsql redis etcd
	err := infrastructure.StartEtcd()
	if err != nil {
		log.Tfatal(conf.Tracer, "Ci StartEtcd Error: %s", err)
	}
	//err := infrastructure.StartConsul()
	//if err != nil {
	//log.Tfatalf(conf.Tracer, "Ci StartConsul Error: %s", err)
	//}
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
	time.Sleep(10 * time.Second)
	test.TestApp()
	//5. Clean App
	err = container.StopApp()
	if err != nil {
		log.Tfatalf(conf.Tracer, "Ci StopApp Error: %s", err)
	}
	err = container.RemoveAppContainer()
	if err != nil {
		log.Tfatalf(conf.Tracer, "Ci CleanContainer Error: %s", err)
	}
	err = infrastructure.StopPostgres()
	if err != nil {
		log.Tfatalf(conf.Tracer, "StopPostgresContainer Error:%s", err)
	}
	err = infrastructure.StopRedis()
	if err != nil {
		log.Tfatal(conf.Tracer, "StopReisContainer Error:%s", err)

	}
	err = infrastructure.StopEtcd()
	if err != nil {
		log.Tfatal(conf.Tracer, "RemoveReisContainer Error:%s", err)
	}
}

func CiPush(traceId, tag string) {
	// 修改镜像Tag --> Push到Repo --> 删除旧镜像
	Registry := conf.Config.Registry
	for _, service := range conf.Config.Services {
		Name := traceId + "-" + service.Name
		Repo := Registry + "/" + service.Name + "/" + Name
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
