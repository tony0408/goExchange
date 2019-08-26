package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/tony0408/goExchange/exchange"
	"github.com/tony0408/goExchange/exchange/stex"
	"github.com/tony0408/goExchange/exchange/stex/model"
)

const STEX = "stex"

func main() {
	flag.Parse()

	if flag.NFlag() < 3 {
		flag.Usage()
		os.Exit(3)
	}

	switch strings.ToLower(exchangeName) {
	case STEX:
		stexExchange := stex.NewStex()
		stexExchange.SetProxy(http.ProxyFromEnvironment)

		switch strings.ToLower(target) {
		case "orderbook":
			for out := range stexExchange.Subscribe(exchange.SubscribeFunc(stex.SubscribeOrderBook), symbols...) {
				fmt.Printf("%+v\n", *(out.(*model.OrderBook)))
			}
		}
	}
}
