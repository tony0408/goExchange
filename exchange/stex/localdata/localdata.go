package localdata

import (
	"github.com/NzKSO/goExchange/exchange/stex/model"
)

const currencyPairsInfoFile = "currencyPairsInfo.json"

// AllCurrencyPairs contain all information related to currency pair in stex
var AllCurrencyPairs map[string]*model.CurrencyPair

func init() {
	/* _, file, _, ok := runtime.Caller(0)
	if !ok {
		panic("no caller information")
	}

	data, err := ioutil.ReadFile(filepath.Join(filepath.Dir(file), currencyPairsInfoFile))
	if err != nil {
		panic(err)
	}

	var allCurrencyPairs []model.CurrencyPair

	if err = json.Unmarshal(data, &allCurrencyPairs); err != nil {
		panic(err)
	} */

	AllCurrencyPairs = make(map[string]*model.CurrencyPair)

	/* for i := range allCurrencyPairs {
		AllCurrencyPairs[allCurrencyPairs[i].Symbol] = &allCurrencyPairs[i]
	} */
}
