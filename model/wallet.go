package model

type Wallet struct {
	Username string  `json:"username"`
	Balance  float64 `json:"balance"`
}
