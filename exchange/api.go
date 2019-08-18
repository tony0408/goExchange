package exchange

import (
	"context"
	"net/http"
	"net/url"
)

// Subscriber is a generic subscribe interface
type Subscriber interface {
	Subscribe(ctx context.Context, identifiers interface{}, out chan interface{})
}

// Exchange defined a group of generic interface for exchange
type Exchange interface {
	Subscribe(sub Subscriber, symbols ...string) <-chan interface{}
	ConvertCurrencyPair(currencyPair string) string
	SetProxy(proxy func(*http.Request) (*url.URL, error))
}
