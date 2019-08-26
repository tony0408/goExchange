package model

// OrderBook represents orderbook pushed through websocket
type OrderBook struct {
	CurrencyPairID int         `json:"currency_pair_id"`
	Amount         string      `json:"amount"`
	Price          string      `json:"price"`
	Amount2        string      `json:"amount2"`
	Count          int         `json:"count"`
	Socket         interface{} `json:"socket"`
}
