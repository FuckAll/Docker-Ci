package test

import (
	"github.com/FuckAll/Docker-Ci/build"
	"github.com/FuckAll/Docker-Ci/conf"
	"github.com/wothing/log"
	"os"
)

func TestApp() {
	command := conf.Config.TestCommand
	if err := os.Chdir(conf.Config.ProjectPath); err != nil {
		log.Terrorf(conf.Tracer, "cd %s Error ", conf.Config.ProjectPath)

	}

	com := build.FMT(command, conf.Tracer)
	_, err := build.CMD(com)
	if err != nil {
		log.Terrorf(conf.Tracer, "Run %s Error", command)
	}
}
