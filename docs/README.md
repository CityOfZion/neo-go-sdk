# Documentation

This package allows you to create a **Client** struct, which can carry out API interactions
on a node in the NEO blockchain.

## Creating a Client

Everything within this package relies on using the **Client** struct.

```golang
nodeURI := "http://test1.cityofzion.io:8880"
client := neo.NewClient(nodeURI)
```

## Checking a Node

The `.NewClient()` function does not check if the node is online, this can be done by using
`.Ping()`:

```golang
ok := client.Ping()
if !ok {
  log.Fatal("Unable to connect to NEO node")
}
```

## API Interactions

The table below covers which API methods are currently supported by this package:

| API Method         | Client Function                 | Parameters      | Return Values        |
|--------------------|---------------------------------|-----------------|----------------------|
| getbestblockhash   | `.GetBestBlockHash()`           | -               | `string, error`      |
| getblock           | `.GetBlockByHash()`             | `string`        | `Block, error`       |
| getblock           | `.GetBlockByIndex()`            | `int64`         | `Block, error`       |
| getblockcount      | `.GetBlockCount()`              | -               | `int64, error`       |
| getblockhash       | `.GetBlockHash()`               | `int64`         | `string, error`      |
| getconnectioncount | `.GetConnectionCount()`         | -               | `int64, error`       |
| getrawtransaction  | `.GetTransaction()`             | `string`        | `Transaction, error` |
| gettxout           | `.GetTransactionOutput()`       | `string, int64` | `Vout, error`        |
| getrawmempool      | `.GetUnconfirmedTransactions()` | -               | `[]string, error`    |

To read more about the NEO node API, please see the [CoZ documentation](https://github.com/CityOfZion/docs/blob/develop/en-us/node/api.md).