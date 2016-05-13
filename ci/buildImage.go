package ci

import (
	"fmt"
	"strings"

	. "github.com/FuckAll/Docker-Ci/cmdrun"
	"github.com/FuckAll/Docker-Ci/conf"
)

func DockerImageBuild() {

	GitLog()
	// plan := make([]string, 0)
	for _, SrvOne := range conf.Services {
		build := fmt.Sprintf(`docker build -t %s:%s %s/%s`, SrvOne.Name, conf.Tracer, conf.ProjectPath, SrvOne.Path)
		build = strings.Replace(build, "\n", "", -1)
		CMD(build)
	}
}
