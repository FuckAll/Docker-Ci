package api

import (
	"github.com/fsouza/go-dockerclient"
)

func CreateContainer(Name, Image string) (string, error) {
	configs := docker.Config{
		Image: Image,
		Tty:   true,
	}
	opts := docker.CreateContainerOptions{
		Name:   Name,
		Config: &configs,
	}
	container, err := client.CreateContainer(opts)
	if err != nil {
		return "", err
	}
	return container.ID, nil
}

func StartContainer(Id, NetworkMode string) error {
	hostConfig := docker.HostConfig{
		NetworkMode: NetworkMode,
	}
	err := client.StartContainer(Id, &hostConfig)
	if err != nil {
		return err
	}
	return nil
}

func StopContainer(id string, timeout uint) error {
	err := client.StopContainer(id, timeout)
	if err != nil {
		return err
	}
	return nil
}

func RemoveContainer(id string, RemoveVolumes bool) error {
	opts := docker.RemoveContainerOptions{
		ID:            id,
		RemoveVolumes: RemoveVolumes,
	}
	err := client.RemoveContainer(opts)
	if err != nil {
		return err
	}
	return nil

}
