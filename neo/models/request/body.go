package request

import "encoding/json"

type (
	// Body is a struct used as the body of a POST HTTP JSON-RPC request.
	Body struct {
		ID         int64         `json:"id"`
		Method     string        `json:"method"`
		Parameters []interface{} `json:"params"`
		Version    string        `json:"jsonrpc"`
	}
)

const (
	apiVersion = "2.0"
)

// NewBody creates a new Body struct, with an empty Params slice.
func NewBody(method string) ([]byte, error) {
	body := Body{
		ID:         1,
		Method:     method,
		Parameters: []interface{}{},
		Version:    apiVersion,
	}

	return json.Marshal(body)
}

// NewBodyWithParameters creates a new Body struct, using the provided parameters slice.
func NewBodyWithParameters(method string, parameters []interface{}) ([]byte, error) {
	body := Body{
		ID:         1,
		Method:     method,
		Parameters: parameters,
		Version:    apiVersion,
	}

	return json.Marshal(body)
}
