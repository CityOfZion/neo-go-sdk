package models

type (
	// Vout holds data about the transaction outputs.
	Vout struct {
		Address string `json:"Address"`
		Asset   string `json:"Asset"`
		N       int    `json:"N"`
		Value   string `json:"Value"`
	}
)
