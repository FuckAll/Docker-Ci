package api

import (
	"bytes"
	"fmt"

	"github.com/fsouza/go-dockerclient"
	"github.com/wothing/log"
)

var endpoint = "tcp://local.docker:2375"

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
	images, err := client.ListImages(docker.ListImagesOptions{All: true})
	if err != nil {
		return err
	}
	for _, img := range images {
		fmt.Println("ID:", img.ID)
		fmt.Println("RepoTags:", img.RepoTags)
	}
	return nil
}

func ExistImage(imageName string) bool {
	images, err := client.ListImages(docker.ListImagesOptions{All: true})
	if err != nil {
		log.Fatal("ExistImage Error:", err)
	}
	for _, image := range images {
		for _, repotage := range image.RepoTags {
			fmt.Println(repotage)
			if imageName == repotage {
				return true
			}
		}

	}
	return false
}

func RemoveImage(name string) error {
	err := client.RemoveImage(name)
	if err != nil {
		return err
	}
	return nil
}

func BuildImage(Name, Dockerfile, ContextDir string, Pull, NoCache, ForceRmTmpContainer bool) error {
	var buf bytes.Buffer
	opts := docker.BuildImageOptions{
		Name:                Name,
		Dockerfile:          Dockerfile,
		OutputStream:        &buf,
		ContextDir:          ContextDir,
		Pull:                Pull,    // Attempt to pull the image even if an older image exists locally.
		NoCache:             NoCache, // Do not use the cache when building the image.
		ForceRmTmpContainer: ForceRmTmpContainer,
	}
	err := client.BuildImage(opts)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
