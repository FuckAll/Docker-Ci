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
	"strings"
	"time"
)

// CiRun Run Ci Begin
func CiRun(step string, tag string, args ...string) {
	Prepare()
	switch step {
	case "OnlyBuild":
		log.Tinfo(conf.Tracer, "OnlyBuild Start!")
		CiBuildApp(args...)
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
		if len(args) < 1 {
			CiPush(conf.Tracer, tag)
		} else {
			CiPush(conf.Tracer, tag, args...)
		}
		log.Tinfo(conf.Tracer, "Push Complate!")
	default:
		fmt.Println("CiRun Do Nothing!!!")
	}
}

// Prepare Used To Create Docker Environment
func Prepare() {
	// Create NetWork Test For Docker Test
	log.Tinfo("Create NetWork Bridge Test")
	bridge := conf.Config.Bridge
	if fi := api.NetworkExist(bridge); !fi {
		_, err := api.CreateNetwork(bridge)
		if err != nil {
			log.Tfatal(conf.Tracer, "Prepare Error: %s", err)
		}
	}
	log.Tinfo("Create Network Bridge Complete!")
	// Prepare Images
	log.Tinfo("Check Infrastructure Images")
	for _, infra := range conf.Config.Infrastructure {
		// image =: reg.17mei.top/redis:latest
		image := (infra.(map[string]interface{})["image"].(string))

		//tmp := []string{"reg.17mei.top","redis:latest"}
		tmp := strings.Split(image, "/")

		// Registry := "reg.17mei.top"
		Registry := tmp[0]

		// ImageTag := []string{"redis","latest"}
		ImageTag := strings.Split((tmp[len(tmp)-1]), ":")

		//Tag := "latest"
		Tag := ImageTag[1]

		//Repository :="reg.17mei.top"
		Repository := Registry

		if len(tmp) > 2 {
			for _, one := range tmp[1 : len(tmp)-1] {
				Repository = Repository + "/" + one

			}

		}

		if err := api.PullImage(Repository, Registry, Tag); err != nil {
			log.Tfatal(conf.Tracer, err)
		}
		log.Tinfo("Check Infrastructure Complete!")
	}

}

// CiBuildApp used to Build App no test
func CiBuildApp(args ...string) {
	//如果OnlyBuild 没有任何的参数就全部Build一遍
	if len(args) < 1 {
		_, err := build.BuildApp()
		if err != nil {
			log.Tfatalf(conf.Tracer, "BuildApp Error: %s", err)
		}
		build.CreateDockerFile()
		build.BuildImage()
	} else {
		// 如果有参数就只Build指定的一些
		newservice := []conf.Service{}
		for _, i := range conf.Config.Services {
			for _, name := range args {
				if name == i.Name {
					newservice = append(newservice, i)
				}
			}
		}
		conf.Config.Services = newservice
		_, err := build.BuildApp()
		if err != nil {
			log.Tfatalf(conf.Tracer, "BuildApp Error: %s", err)
		}
		build.CreateDockerFile()
		build.BuildImage()
	}

}

// CiTestAppNoClean Used To Test App And No Clean Docker Images
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
	time.Sleep(5 * time.Second)
	//3. 启动业务代码容器
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
	time.Sleep(5 * time.Second)
	//3. 启动业务代码容器
	err = container.StartApp()
	if err != nil {
		log.Tfatalf(conf.Tracer, "Ci StartApp Error: %s", err)
	}
	//4. 测试
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

func CiPush(traceId, tag string, apps ...string) {
	// 修改镜像Tag --> Push到Repo --> 删除旧镜像
	Registry := conf.Config.Registry
	newservice := []conf.Service{}
	if len(apps) > 0 {
		for _, i := range conf.Config.Services {
			for _, name := range apps {
				if name == i.Name {
					newservice = append(newservice, i)
				}
			}
		}

	} else {
		newservice = conf.Config.Services
	}
	for _, service := range newservice {
		Name := traceId + "-" + service.Name
		Repo := Registry + "/" + service.Name + "/" + Name
		Tag := tag
		err := api.ChangeTag(Repo, Tag, Name)
		if err != nil {
			log.Tfatalf(conf.Tracer, "Ci CiPush ChangeTag  Error: %v", err)
		}
		err = api.PushImage(Repo, Tag, Registry)
		if err != nil {
			log.Tfatalf(conf.Tracer, "Ci CiPush PushImage  Error: %v", err)
		}
		err = api.RemoveImage(Name)
		if err != nil {
			log.Tfatalf(conf.Tracer, "Ci CiPush RemoveImage Error: %v", err)
		}
	}
}
