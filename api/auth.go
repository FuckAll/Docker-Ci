package api

import (
	"github.com/fsouza/go-dockerclient"
)

func AuthFromDockercfg() (*docker.AuthConfigurations, error) {
	auth, err := docker.NewAuthConfigurationsFromDockerCfg()
	if err != nil {
		return auth, err
	}
	return auth, nil
}
