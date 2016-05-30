package test

import (
	"github.com/FuckAll/Docker-Ci/build"
	"github.com/FuckAll/Docker-Ci/conf"
)

func TestApp() error {
	conf.Tracer = "c02abfe2"
	//command := conf.Config.TestCommand
	//if err := os.Chdir(conf.Config.ProjectPath); err != nil {
	//log.Tfatalf(conf.Tracer, "cd %s Error ", conf.Config.ProjectPath)

	//}

	//_, err := build.CMD(build.FMT(command, conf.Tracer))
	//if err != nil {
	//log.Tfatalf(conf.Tracer, "Run %s Error", command)
	//}
	//log.Tinfof(conf.Tracer, "TestComplate!")
	_, err := build.CMD(build.FMT("CGO_ENABLED=0 go test -c -o /app/testbin %s/gateway/tests/*.go", conf.Config.ProjectPath))
	if err != nil {
		return err
	}
	_, err = build.CMD(build.FMT("TestEnv=CI CiTracer=%s /app/testbin -test.v ", conf.Tracer))
	if err != nil {
		return err
	}
	return nil

}
