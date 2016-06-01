package test

import (
	"fmt"

	"github.com/FuckAll/Docker-Ci/build"
	"github.com/FuckAll/Docker-Ci/conf"
	"github.com/wothing/log"
	"os"
)

func TestApp() {
	conf.Tracer = "de5110f2"
	command := conf.Config.TestCommand
	if err := os.Chdir(conf.Config.ProjectPath); err != nil {
		log.Tfatalf(conf.Tracer, "cd %s Error ", conf.Config.ProjectPath)

	}

	com := build.FMT(command, conf.Tracer)
	fmt.Println(com)
	_, err := build.CMD(com)
	if err != nil {
		log.Tfatalf(conf.Tracer, "Run %s Error", command)
	}
	//log.Tinfof(conf.Tracer, "TestComplate!")
	//_, err := build.CMD(build.FMT("CGO_ENABLED=0 go test -c -o /app/testbin %s/gateway/tests/*.go", conf.Config.ProjectPath))
	//if err != nil {
	//return err
	//}
	//_, err = build.CMD(build.FMT("TestEnv=CI CiTracer=%s /app/testbin -test.v ", conf.Tracer))
	//if err != nil {
	//return err
	//}

}
