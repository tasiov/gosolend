package models

type Reserve struct {
	Address         string  `json:"address"`
	LiquidityAmount uint64  `json:"liquidityAmount"`
	BorrowRate      float64 `json:"borrowRate"`
	LendingRate     float64 `json:"lendingRate"`
	LastUpdated     int64   `json:"lastUpdated"`
}
