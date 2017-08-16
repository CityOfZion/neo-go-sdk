<p align="center">
  <img 
    src="https://res.cloudinary.com/vidsy/image/upload/v1501603448/logo_w2erl5.jpg" 
    width="100px"
  >
</p>

<h1 align="center">neo-go-sdk</h1>

<p align="center">
  Golang SDK for the <b>NEO</b> blockchain.
</p>

<p align="center">
  <a href="https://github.com/CityOfZion/neo-go-sdk/releases">
    <img src="https://img.shields.io/github/tag/CityOfZion/neo-go-sdk.svg?style=flat">
  </a>
  <a href="https://circleci.com/gh/CityOfZion/neo-go-sdk/tree/master">
    <img src="https://circleci.com/gh/CityOfZion/neo-go-sdk/tree/master.svg?style=shield">
  </a>
  <a href="https://goreportcard.com/report/github.com/CityOfZion/neo-go-sdk">
    <img src="https://goreportcard.com/badge/github.com/CityOfZion/neo-go-sdk">
  </a>
</p>

## What?

- Client for interacting with a node on the [NEO](http://neo.org/) blockchain.
- Retrieve data and send actions.
- Fully tested [Golang](https://golang.org/) package.
- Aimed to help other developers build applications for the NEO ecosystem.
- Written using the standard library, without 3rd party packages. 

## Quick Start

```
go get github.com/CityOfZion/neo-go-sdk
```

```golang
package main

import (
  "log"

  "github.com/CityOfZion/neo-go-sdk/neo"
)

func main() {
  nodeURI := "http://test1.cityofzion.io:8880"
  client := neo.NewClient(nodeURI)

  ok := client.Ping()
  if !ok {
    log.Fatal("Unable to connect to NEO node")
  }

  block, err := client.GetBlockByHash(
    "3f0b498c0d57f73c674a1e28045f5e9a0991f9dac214076fadb5e6bafd546170",
  )
  if err != nil {
    log.Fatal(err)
  }

  log.Printf("Block found, index is: %d", block.Index)
}
```

## Examples

Full documentation of how to use this package can be found: [docs/README.md](https://github.com/CityOfZion/neo-go-sdk/blob/master/docs)

## Help

- Open a new [issue](https://github.com/CityOfZion/neo-go-sdk/issues/new) if you encountered a problem.
- Or ping **@revett** on the [NEO Slack](https://join.slack.com/t/neoblockchainteam/shared_invite/MjE3ODMxNDUzMDE1LTE1MDA4OTY3NDQtNTMwM2MyMTc2NA).
- Submitting PRs to the project is always welcome! 🎉
- Check the [Changelog](https://github.com/CityOfZion/neo-go-sdk/blob/master/CHANGELOG.md) for recent changes.

## License

- Open-source [MIT](https://github.com/CityOfZion/neo-go-sdk/blob/master/LICENSE) 