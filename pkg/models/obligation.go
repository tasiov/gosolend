package models

type Obligation struct {
	Owner       string   `json:"owner"`
	Deposits    []uint64 `json:"deposits"`
	Borrows     []uint64 `json:"borrows"`
	LastUpdated int64    `json:"lastUpdated"`
}
