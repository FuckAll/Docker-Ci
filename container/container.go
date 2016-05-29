package container

import (
	//	"bytes"
	"errors"
	//	"fmt"
	//	"os"
	//	"os/exec"
	//	"time"

	"github.com/FuckAll/Docker-Ci/api"
	"github.com/FuckAll/Docker-Ci/conf"
)

var expire = 20 // stop container expire time

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
		//containerName := conf.Tracer + "-" + service
		containerName := conf.Tracer + "-" + service.Name
		imageName := conf.Tracer + "-" + service.Name + ":latest"
		//if or not container exist
		var env []string
		for k, v := range service.Env {
			switch k {
			case "CH", "RH", "DH":
				tmp := k + "=" + conf.Tracer + "-" + v
			default:
				tmp := k + "=" + v.(string)
			}
			env = append(env, tmp)
		}
		if !api.ExistImage(imageName) {
			return []string{""}, errors.New("CreateImages ExistImage Error")
		}
		containerId, err := api.CreateContainer(containerName, imageName, []string{"app:/test"}, env...)
		if err != nil {
			return []string{""}, err
		}
		containerIds = append(containerIds, containerId)
	}
	AppContainerIds = containerId
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

func RemoveAppContrainer() error {
	for _, containerId := range AppContainerIds {
		err := api.RemoveContainer(containerId, false)
		if err != nil {
			return err
		}
	}
	return nil
}
