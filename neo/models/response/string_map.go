package response

type (
	// StringMap represents the JSON schema of a response from a NEO node, where the
	// expected result is an object where the keys are strings.
	StringMap struct {
		ID      int                    `json:"id"`
		JSONRPC string                 `json:"jsonrpc"`
		Result  map[string]interface{} `json:"result"`
	}
)
