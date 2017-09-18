<p align="center">
  <img 
    src="http://res.cloudinary.com/vidsy/image/upload/v1503160820/CoZ_Icon_DARKBLUE_200x178px_oq0gxm.png" 
    width="125px"
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
  <a href="https://godoc.org/github.com/CityOfZion/neo-go-sdk/neo">
    <img src="https://godoc.org/github.com/CityOfZion/neo-go-sdk?status.svg">
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

See [GoDoc](https://godoc.org/github.com/CityOfZion/neo-go-sdk/neo) for full documentation.

## CLI

Debugging a NEO public and private key pair is a common task when interacting with the blockchain.
Make use of the **neo-go-sdk** CLI to help with this process:

```
./neo-go-sdk --wif KxQREAjBL6Ga8dw9rPN45pwoZ5dxhAQacEajQke6qmpB7DW6nAWE
```

This will output the **full details** about the key pair. See [releases](https://github.com/CityOfZion/neo-go-sdk/releases) to download the CLI.

## Help

- Open a new [issue](https://github.com/CityOfZion/neo-go-sdk/issues/new) if you encountered a problem.
- Or ping **@revett** on the [NEO Slack](https://neo-slack-invite.herokuapp.com).
- Submitting PRs to the project is always welcome! 🎉
- Check the [Changelog](https://github.com/CityOfZion/neo-go-sdk/blob/master/CHANGELOG.md) for recent changes.

## License

- Open-source [MIT](https://github.com/CityOfZion/neo-go-sdk/blob/master/LICENSE).
- Main author is [@revett](https://github.com/revett).
- This project adheres to the [Contributor Covenant Code of Conduct](https://github.com/goreleaser/goreleaser/blob/master/CODE_OF_CONDUCT.md).