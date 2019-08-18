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

// CurrencyPair contains some information about currency pair
type CurrencyPair struct {
	ID                int    `json:"id"`
	CurrencyID        int    `json:"currency_id"`
	CurrencyCode      string `json:"currency_code"`
	CurrencyName      string `json:"currency_name"`
	MarketCurrencyID  int    `json:"market_currency_id"`
	MarketCode        string `json:"market_code"`
	MarketName        string `json:"market_name"`
	MinOrderAmount    string `json:"min_order_amount"`
	MinBuyPrice       string `json:"min_buy_price"`
	MinSellPrice      string `json:"min_sell_price"`
	BuyFeePercent     string `json:"buy_fee_percent"`
	SellFeePercent    string `json:"sell_fee_percent"`
	Active            bool   `json:"active"`
	Delisted          bool   `json:"delisted"`
	PairMessage       string `json:"pair_message"`
	CurrencyPrecision int    `json:"currency_precision"`
	MarketPrecision   int    `json:"market_precision"`
	Symbol            string `json:"symbol"`
	GroupName         string `json:"group_name"`
	GroupID           int    `json:"group_id"`
	AmountMultiplier  int    `json:"amount_multiplier"`
}
