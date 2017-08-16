package models

type (
	// Script holds all data about the scripts used.
	Script struct {
		Invocation   string `json:"Invocation"`
		Verification string `json:"Verification"`
	}
)
