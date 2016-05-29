/**
 * Copyright 2015-2016, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by izgnod on 2016/05/13 22:10
 */

package ci

import (
	"fmt"

	"github.com/FuckAll/Docker-Ci/build"
	"github.com/FuckAll/Docker-Ci/conf"
	"github.com/FuckAll/Docker-Ci/container"
	"github.com/FuckAll/Docker-Ci/infrastructure"
	"github.com/wothing/log"
)

func CiRun() {
	//1.编译代码，制作docker 镜像
	_, err := build.BuildApp()
	if err != nil {
		log.Tfatalf(conf.Tracer, "BuildApp Error: %s", err)
	}
	build.CreateDockerFile()
	build.BuildImage()

	//2. 启动基础服务，例如：pgsql redis consul
	//conf.Tracer = "b29e1b88"
	err = infrastructure.StartConsul()
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
	//5. 根据版本选择是否清除环境

	fmt.Println("CiRun Complate!")

}
