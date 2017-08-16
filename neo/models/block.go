package models

type (
	// Block holds all the data about a particular block on the blockchain.
	Block struct {
		Confirmations     int64         `json:"Confirmations"`
		Hash              string        `json:"Hash"`
		Index             int64         `json:"Index"`
		Merkleroot        string        `json:"Merkleroot"`
		NextBlockHash     string        `json:"Nextblockhash"`
		NextConsensus     string        `json:"Nextconsensus"`
		Nonce             string        `json:"Nonce"`
		PreviousBlockHash string        `json:"Previousblockhash"`
		Size              int64         `json:"Size"`
		Time              int64         `json:"Time"`
		Version           int64         `json:"Version"`
		Script            Script        `json:"Script"`
		Transactions      []Transaction `json:"Tx"`
	}
)
