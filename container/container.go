package container

import (
	"errors"
	"github.com/FuckAll/Docker-Ci/api"
	"github.com/FuckAll/Docker-Ci/conf"
)

var expire uint = 20 // stop container expire time

var AppContainerIds []string

func StartApp() error {
	err := CreateAppContainer()
	if err != nil {
		return err
	}
	err = StrartAppContainer()
	if err != nil {
		return err
	}
	return nil
}

func StopApp() error {
	err := StopAppContainer()
	if err != nil {
		return err
	}
	return nil

}

func CreateAppContainer() error {
	var containerIds []string
	for _, service := range conf.Config.Services {
		var env []string
		containerName := conf.Tracer + "-" + service.Name
		imageName := conf.Tracer + "-" + service.Name + ":latest"
		for k, v := range service.Env {
			var tmp string
			switch k {
			case "ETCD", "SVC_REDIS_SERVICE_HOST", "SVC_PGSQL_SERVICE_HOST":
				tmp = k + "=" + conf.Tracer + "-" + v.(string)
			default:
				tmp = k + "=" + v.(string)
			}
			env = append(env, tmp)
		}
		if !api.ExistImage(imageName) {
			return errors.New("CreateImages ExistImage Error")
		}
		containerId, err := api.CreateContainer(containerName, imageName, []string{"app:/test"}, env...)
		if err != nil {
			return err
		}
		containerIds = append(containerIds, containerId)
	}
	AppContainerIds = containerIds
	return nil
}

func StrartAppContainer() error {
	for _, containerId := range AppContainerIds {
		err := api.StartContainer(containerId, conf.Config.Bridge)
		if err != nil {
			return err
		}
	}
	return nil

}

func StopAppContainer() error {
	for _, containerId := range AppContainerIds {
		err := api.StopContainer(containerId, expire)
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveAppContainer() error {
	for _, containerId := range AppContainerIds {
		err := api.RemoveContainer(containerId, false)
		if err != nil {
			return err
		}
	}
	return nil
}
