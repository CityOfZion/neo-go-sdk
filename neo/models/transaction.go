package models

type (
	// Transaction holds all data about a transaction on the blockchain.
	Transaction struct {
		ID         string        `json:"Txid"`
		Size       int64         `json:"Size"`
		Type       string        `json:"Type"`
		Version    int64         `json:"Version"`
		Attributes []interface{} `json:"Attributes"` // 💩
		Vin        []Vin         `json:"Vin"`
		Vout       []Vout        `json:"Vout"`
		SysFee     string        `json:"Sys_fee"`
		NetFee     string        `json:"Net_fee"`
		Scripts    []Script      `json:"Scripts"`
		Nonce      int64         `json:"Nonce"`
	}
)
