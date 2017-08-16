package neo_test

import (
	"fmt"

	"github.com/CityOfZion/neo-go-sdk/neo"
)

const (
	cozNodeURI = "http://test%d.cityofzion.io:8880"
)

func selectTestNode() string {
	var nodeURI string

	for i := 1; i <= 5; i++ {
		uri := fmt.Sprintf(cozNodeURI, i)
		client := neo.NewClient(uri)

		ok := client.Ping()
		if ok {
			nodeURI = uri
			break
		}
	}

	if nodeURI == "" {
		panic("No available nodes for testing.")
	}

	return nodeURI
}
