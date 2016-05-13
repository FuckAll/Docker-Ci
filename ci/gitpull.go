package ci

import (
	. "github.com/FuckAll/Docker-Ci/cmdrun"
	"github.com/FuckAll/Docker-Ci/conf"
	"github.com/wothing/log"
)

func GitPull() {
	CMD(FMT("cd %s && git pull %s", conf.ProjectPath, conf.REPO))
	GitLog()
	log.Tinfo(conf.Tracer, FMT("REPO: %s pull done", conf.REPO))
}

func GitLog() {
	conf.Tracer = CMD(FMT("cd %s && git rev-parse HEAD", conf.ProjectPath))
}
