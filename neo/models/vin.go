package models

type (
	// Vin is holds data about the transaction input.
	Vin struct {
		TransactionID string `json:"Txid"`
		Vout          int    `json:"Vout"`
	}
)
