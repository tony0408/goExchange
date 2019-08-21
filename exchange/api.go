package exchange

import (
	"context"
	"net/http"
	"net/url"

	"github.com/bitontop/gored/exchange/stex"
)

// ContextKey represents a custom type of context key, which can be used in context.WithValue
type ContextKey string

// SubscribeFunc represents type of subscription function
type SubscribeFunc func(context.Context, interface{}, chan<- interface{})

// Subscribe implements interface Subscriber
func (f SubscribeFunc) Subscribe(ex *stex.Stex, ctx context.Context, identifier interface{}, ch chan interface{}) {
	f(ex, ctx, identifier, ch)
}

// Subscriber is a generic subscribe interface
type Subscriber interface {
	Subscribe(ex *stex.Stex, ctx context.Context, identifiers interface{}, out chan interface{})
}

// Exchange defined a group of generic interface for exchange
type Exchange interface {
	Subscribe(ex *stex.Stex, sub Subscriber, symbols ...string) <-chan interface{}
	ConvertCurrencyPair(currencyPair string) string
	SetProxy(proxy func(*http.Request) (*url.URL, error))
}
