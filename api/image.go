package api

import (
	"bytes"
	"fmt"
	"time"

	"github.com/fsouza/go-dockerclient"
	"github.com/wothing/log"
)

var endpoint = "tcp://192.168.64.3:2375"

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

func PushImage(Name, Tag, Registry string) error {
	var buf bytes.Buffer
	opts := docker.PushImageOptions{
		Name:              Name,
		Tag:               Tag,
		Registry:          Registry,
		OutputStream:      &buf,
		RawJSONStream:     false,
		InactivityTimeout: time.Second * 100,
	}
	auth, err := AuthFromDockercfg()
	fmt.Println(opts)
	if err != nil {
		return err
	}
	err = client.PushImage(opts, auth.Configs[Registry])
	if err != nil {
		return err
	}
	return nil
}

func ChangeTag(Repo, Tag, Name string) {
	err := client.TagImage
}
