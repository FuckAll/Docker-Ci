package dockerapi

import (
	"github.com/fsouza/go-dockerclient"
	"github.com/wothing/log"
)

func AuthFromDockercfg() (*docker.AuthConfigurations, error) {
	auth, err := docker.NewAuthConfigurationsFromDockerCfg()
	if err != nil {
		log.Terror("1111", "AuthFromDockercfg Error")
		return auth, err
	}
	return auth, nil
}
