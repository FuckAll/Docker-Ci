package api

import (
	"github.com/fsouza/go-dockerclient"
)

func ListNetwork() ([]docker.Network, error) {
	networks, err := client.ListNetworks()
	if err != nil {
		return nil, err
	}
	return networks, nil

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
		return err
	}
	return nil
}

func NetworkExist(Name string) bool {
	networks, err := ListNetwork()
	if err != nil {
		return false
	}
	for _, net := range networks {
		if net.Name == Name {
			return true
		}
	}
	return false
}
