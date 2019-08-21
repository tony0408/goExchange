package main

import (
	"flag"
	"fmt"
	"strings"
)

type symbolsFlag []string

var (
	symbols      symbolsFlag
	exchangeName string
	target       string
)

const (
	symbolsUsage  = "The stock symbols to subscribe, it is formed of currency/market, which accepts multiple currency/market splitted by comma, ETH/BTC,BTC/USD for example"
	exchangeUsage = "Specify exchange name, currently supports stex only"
	targetUsage   = "Specify which data to get from exchange, currently supports orderbook only"
)

func (f *symbolsFlag) String() string {
	return strings.Join(*f, ",")
}

func (f *symbolsFlag) Set(s string) error {
	*f = strings.Split(s, ",")

	for i, v := range *f {
		if subStrs := strings.Split(v, "_"); len(subStrs) != 2 {
			return fmt.Errorf("invalid symbol %q", v)
		}
		(*f)[i] = strings.TrimSpace(v)
	}
	return nil
}

func init() {
	flag.Var(&symbols, "s", symbolsUsage)
	flag.StringVar(&exchangeName, "e", "", exchangeUsage)
	flag.StringVar(&target, "t", "", targetUsage)
}
