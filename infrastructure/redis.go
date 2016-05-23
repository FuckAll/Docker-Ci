package infrastructure

import (
	"fmt"
	"github.com/FuckAll/Docker-Ci/api"
)

type CreateRedisOpts struct {
	Image  string
	Name   string
	Port   string
	Passwd string
	Init   string
}

func CreateRedis(opts *CreateRedisOpts) (string, error) {
	name := "redis-" + opts.Name
	containerId, err := api.CreateContainer(name, opts.Image, []string{})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return containerId, nil
}

func StartRedis(id string, networkmode string) error {
	err := api.StartContainer(id, networkmode)
	if err != nil {
		fmt.Println(err)
		return err

	}
	return nil
}

func RemoveRedis(id string, removevolemes bool) error {
	err := api.RemoveContainer(id, removevolemes)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
