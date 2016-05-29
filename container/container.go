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

func CreateAppContainer() ([]string, error) {
	var containerIds []string
	for _, service := range conf.Config.Services {
		//containerName := conf.Tracer + "-" + service
		containerName := conf.Tracer + "-" + service.Name
		imageName := conf.Tracer + "-" + service.Name + ":latest"
		//if or not container exist
		var env []string
		for k, v := range service.Env {
			tmp := k + "=" + v.(string)
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
	return containerIds, nil
}

func StrartAppContainer(containerIds []string) error {
	for _, containerId := range containerIds {
		err := api.StartContainer(containerId, conf.Config.Bridge)
		if err != nil {
			return err
		}
	}
	return nil

}

func StopAppContainer(containerIds []string) error {
	for _, containerId := range containerIds {
		err := api.StopContainer(containerId, expire)
		if err != nil {
			return err
		}
	}
	return nil
}

func RemoveAppContrainer(containerIds []string) error {
	for _, containerId := range containerIds {
		err := api.RemoveContainer(containerId, false)
		if err != nil {
			return err
		}
	}
	return nil
}
