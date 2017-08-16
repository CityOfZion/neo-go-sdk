package response

type (
	// StringArray represents the JSON schema of a response from a NEO node, where the
	// expected result is an array of strings.
	StringArray struct {
		ID      int      `json:"id"`
		JSONRPC string   `json:"jsonrpc"`
		Result  []string `json:"result"`
	}
)
