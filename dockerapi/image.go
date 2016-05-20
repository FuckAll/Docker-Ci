package dockerapi

import (
	"bytes"
	"fmt"

	"github.com/fsouza/go-dockerclient"
)

var endpoint = "tcp://42.96.131.96:2375"

var client *docker.Client

func init() {
	var err error
	client, err = docker.NewClient(endpoint)
	if err != nil {
		fmt.Println(err)
	}
}
func PullImage(Repository, Registry, Tag string) error {
	auth, err := AuthFromDockercfg()
	if err != nil {
		return err
	}
	fmt.Println(auth.Configs[Registry])
	err = client.PullImage(docker.PullImageOptions{Repository: Repository, Registry: Registry, Tag: Tag}, auth.Configs[Registry])
	if err != nil {
		return err
	}
	return nil
}

func ListImages(All bool) error {
	images, err := client.ListImages(docker.ListImagesOptions{All: All})
	if err != nil {
		return err
	}
	for _, img := range images {
		fmt.Println("ID:", img.ID)
		fmt.Println("RepoTags:", img.RepoTags)
	}
	return nil
}

func RemoveImage(name string) error {
	err := client.RemoveImage(name)
	if err != nil {
		return err
	}
	return nil
}

func BuildImage(Name, Remote string, NoCache bool) error {
	var buf bytes.Buffer
	opts := docker.BuildImageOptions{
		Name:           Name,
		Remote:         Remote,
		SuppressOutput: true,
		OutputStream:   &buf,
	}
	err := client.BuildImage(opts)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
