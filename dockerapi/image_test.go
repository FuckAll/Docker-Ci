package dockerapi

import (
	"fmt"
	"testing"
)

func TestPullImage(t *testing.T) {
	Repository := "index.tenxcloud.com/izgnod/consul"
	Registry := "index.tenxcloud.com"
	Tag := "v1"
	err := PullImage(Repository, Registry, Tag)
	if err != nil {
		fmt.Println("TestPullImage Error")
	}
}

func TestListImage(t *testing.T) {
	if ListImages(false) != nil {
		fmt.Println("TestListImage Error")
	}
}

func TestBuildImage(t *testing.T) {
	err := BuildImage("Test", "Dockerfile", false)
	//err := BuildImage("Test", "/root/Dockerfile", false)
	if err != nil {
		fmt.Println("TestBuildImage Error")
	}
}
