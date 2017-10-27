package neo_test

import (
	"fmt"

	"github.com/CityOfZion/neo-go-sdk/neo"
)

const (
	neoNodeURI = "http://seed%d.neo.org:10332"
)

func selectTestNode() string {
	var nodeURI string

	for i := 1; i <= 5; i++ {
		uri := fmt.Sprintf(neoNodeURI, i)
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
