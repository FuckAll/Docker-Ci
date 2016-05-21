package api

import (
	"fmt"
	"github.com/fsouza/go-dockerclient"
)

func ListNetwork() error {
	networks, err := client.ListNetworks()
	if err != nil {
		return err
	}
	//for _, id := range networks {
	//fmt.Println(id.ID)
	//fmt.Println(id.Containers)
	//}
	fmt.Println(networks)
	return nil

}

func CreateNetwork(Name string) (string, error) {
	NetworkOptions := docker.CreateNetworkOptions{
		Name: Name,
	}

	network, err := client.CreateNetwork(NetworkOptions)
	if err != nil {
		return "", err
	}
	return network.Name, nil
}

func ConnectNetwork(Id string) error {
	err := client.ConnectNetwork(Id, docker.NetworkConnectionOptions{})
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
