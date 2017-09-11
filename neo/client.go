package neo

import (
	"net"
	"net/url"

	"github.com/CityOfZion/neo-go-sdk/neo/models"
	"github.com/CityOfZion/neo-go-sdk/neo/models/response"
)

type (
	// Client is the entrypoint for the package, it is used to carry out all actions.
	Client struct {
		NodeURI string
	}
)

// NewClient creates a new Client struct using the providing NodeURI argument.
func NewClient(nodeURI string) Client {
	return Client{
		NodeURI: nodeURI,
	}
}

// GetBestBlockHash returns the hash of the best block in the chain.
func (c Client) GetBestBlockHash() (string, error) {
	var response response.String

	err := executeRequest("getbestblockhash", nil, c.NodeURI, &response)
	if err != nil {
		return "", err
	}

	return response.Result, nil
}

// GetBlockByHash returns the corresponding block information according to the specified
// hash value.
func (c Client) GetBlockByHash(hash string) (*models.Block, error) {
	requestBodyParams := []interface{}{
		hash, 1,
	}
	var response response.Block

	err := executeRequest("getblock", requestBodyParams, c.NodeURI, &response)
	if err != nil {
		return nil, err
	}

	return &response.Result, nil
}

// GetBlockByIndex returns the corresponding block information according to the specified
// index value.
func (c Client) GetBlockByIndex(index int64) (*models.Block, error) {
	requestBodyParams := []interface{}{
		index, 1,
	}
	var response response.Block

	err := executeRequest("getblock", requestBodyParams, c.NodeURI, &response)
	if err != nil {
		return nil, err
	}

	return &response.Result, nil
}

// GetBlockCount returns the number of blocks in the chain.
func (c Client) GetBlockCount() (int64, error) {
	var response response.Integer

	err := executeRequest("getblockcount", nil, c.NodeURI, &response)
	if err != nil {
		return 0, err
	}

	return response.Result, nil
}

// GetBlockHash returns the hash value of the corresponding block based on the specified
// index.
func (c Client) GetBlockHash(index int64) (string, error) {
	requestBodyParams := []interface{}{
		index,
	}
	var response response.String

	err := executeRequest("getblockhash", requestBodyParams, c.NodeURI, &response)
	if err != nil {
		return "", err
	}

	return response.Result, nil
}

// GetConnectionCount returns the current number of connections for the node.
func (c Client) GetConnectionCount() (int64, error) {
	var response response.Integer

	err := executeRequest("getconnectioncount", nil, c.NodeURI, &response)
	if err != nil {
		return 0, err
	}

	return response.Result, nil
}

// GetTransaction returns the corresponding transaction information based on the
// specified hash value.
func (c Client) GetTransaction(hash string) (*models.Transaction, error) {
	requestBodyParams := []interface{}{
		hash, 1,
	}
	var response response.Transaction

	err := executeRequest("getrawtransaction", requestBodyParams, c.NodeURI, &response)
	if err != nil {
		return nil, err
	}

	return &response.Result, nil
}

// GetTransactionOutput returns the corresponding transaction output (change) information
// based on the specified hash and index.
func (c Client) GetTransactionOutput(hash string, index int64) (*models.Vout, error) {
	requestBodyParams := []interface{}{
		hash, index,
	}
	var response response.Vout

	err := executeRequest("gettxout", requestBodyParams, c.NodeURI, &response)
	if err != nil {
		return nil, err
	}

	return &response.Result, nil
}

// GetUnconfirmedTransactions returns a slice of transaction hashes that are all
// unconfirmed transactions that the node has in memory.
func (c Client) GetUnconfirmedTransactions() ([]string, error) {
	var response response.StringArray

	err := executeRequest("getrawmempool", nil, c.NodeURI, &response)
	if err != nil {
		return nil, err
	}

	return response.Result, nil
}

// Ping checks if the node is online.
func (c Client) Ping() bool {
	parsedURI, err := url.Parse(c.NodeURI)
	if err != nil {
		return false
	}

	conn, err := net.Dial("tcp", parsedURI.Host)
	if err != nil {
		return false
	}

	_ = conn.Close()

	return true
}
