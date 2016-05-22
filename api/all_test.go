package api

import (
	"fmt"
	"testing"
)

var Id string

// List Image
//func TestListImage(t *testing.T) {
//if ListImages(false) != nil {
//fmt.Println("TestListImage Error")
//}
//}

// Pull Image
//func TestPullImage(t *testing.T) {
//Repository := "index.tenxcloud.com/izgnod/consul"
//Registry := "index.tenxcloud.com"
//Tag := "v1"
//err := PullImage(Repository, Registry, Tag)
//if err != nil {
//fmt.Println("TestPullImage Error")
//}
//}

// Build Image
//func TestBuildImage(t *testing.T) {
//err := BuildImage("test", "Dockerfile", "/Users/KongFu/BaseEnv/src/github.com/FuckAll/Docker-Ci/dockerapi", false, false, false)
//if err != nil {
//fmt.Println("TestBuildImage Error")
//}
//}

// Remove Image
//func TestRemoveImage(t *testing.T) {
//err := RemoveImage("test")
//if err != nil {
//fmt.Println("TestRemoveImage Error")
//}
//}

// List Network
//func TestListNetwork(t *testing.T) {

//err := ListNetwork()
//if err != nil {
//fmt.Println("ListNetwork Error")
//fmt.Println(err)
//}
//fmt.Println("ListNetwork Complate")
//}

//Create Network
func TestCreateNetwork(t *testing.T) {
	if fi := NetworkExist("test"); !fi {
		name, err := CreateNetwork("test")
		if err != nil {
			fmt.Println("TestCreateNetwork Error")
			fmt.Println(err)
		}
		fmt.Println("CreateNetwork Name:", name)
	} else {
		fmt.Println("CreateNetwork Complate")
	}
}

// Network Exist
func TestNetworkExist(t *testing.T) {
	fi := NetworkExist("test")
	if !fi {
		fmt.Println("TestNetworkExist Error")
		fmt.Println(fi)
	} else {
		fmt.Println("TestNetworkExist Complate")

	}

}

// Create Container
func TestCreateContainer(t *testing.T) {
	var err error
	Id, err = CreateContainer("test", "test")
	if err != nil {
		fmt.Println("TestCreateContainer Error")
		fmt.Println(err)
	}
	fmt.Println("CreateContainer ID:", Id)
}

// ConnectNetWork
//func TestConnectNetwork(t *testing.T) {
//err := ConnectNetwork(Id)
//if err != nil {
//fmt.Println("TestConnectNetwork Error")
//fmt.Println(err)
//}
//fmt.Println("TestConnectNetwork Complate")

//}

// Start Container
func TestStartContainer(t *testing.T) {
	err := StartContainer(Id, "test")
	if err != nil {
		fmt.Println("TestStartContainer Error")
		fmt.Println(err)
	}
	fmt.Println("StartContainer ID:", Id)
}

// Stop Container
func TestStopContainer(t *testing.T) {
	err := StopContainer(Id, 0)
	if err != nil {
		fmt.Println("TestStopContainer Error")
		fmt.Println(err)
	}
	fmt.Println("TestStopContainer ID:", Id)
}

//Remove Container
func TestRemoveContainer(t *testing.T) {
	err := RemoveContainer(Id, false)
	if err != nil {
		fmt.Println("TestRemoveContainer Error")
		fmt.Println(err)
	}
	fmt.Println("TestRemoveContainer ID:", Id)

}
