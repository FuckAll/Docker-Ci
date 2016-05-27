package build

import (
	"fmt"
	"os"

	"github.com/FuckAll/Docker-Ci/api"
	"github.com/FuckAll/Docker-Ci/conf"
	"github.com/wothing/log"
)

func BuildImage() {
	services := conf.Config.Services
	for _, service := range services {
		imageName := conf.Tracer + "-" + service.Name
		filename := service.Name + "-DockerFile"
		err := api.BuildImage(imageName, filename, GoPath, false, false, false)
		if err != nil {
			fmt.Println(err)
		}

	}

}

func CreateDockerFile() {
	if err := os.Chdir(GoPath); err != nil {
		log.Tfatalf(conf.Tracer, "cd %s Error ", GoPath)
	}
	for _, service := range conf.Config.Services {
		destanition := service.DockerFile["CopyTo"]
		cmd := service.DockerFile["CMD"]

		filepath := GoPath + `/` + service.Name + "-DockerFile"
		file, err := os.Create(filepath)
		if err != nil {
			log.Tfatal(conf.Tracer, err)
		}
		file.WriteString(fmt.Sprintln(`FROM` + ` ` + conf.Config.ServicesImage))
		file.WriteString(fmt.Sprintln(`COPY ` + service.Name + ` ` + destanition.(string)))
		file.WriteString(`CMD ` + string(cmd.(string)))
		file.Close()
	}

}
