package api

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
)

func ListVolumes(opts docker.ListVolumesOptions) ([]docker.Volume, error) {
	volumes, err := client.ListVolumes(opts)
	if err != nil {
		fmt.Println(err)
		return []docker.Volume{}, err
	}
	return volumes, err
}

func CreateVolume(name string) (string, error) {
	opts := docker.CreateVolumeOptions{
		Name: name,
	}
	volume, err := client.CreateVolume(opts)
	if err != nil {
		fmt.Println(err)
		return "", nil

	}
	return volume.Name, err
}

func RemoveVolume(name string) error {
	err := client.RemoveVolume(name)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func VolumeExist(name string) bool {
	volumes, err := ListVolumes(docker.ListVolumesOptions{})
	if err != nil {
		fmt.Println(err)
		return false
	}
	for _, v := range volumes {
		if v.Name == name {
			return true
		}
	}
	return false
}
