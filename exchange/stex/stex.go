package stex

import (
	"context"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/NzKSO/goExchange/exchange"
	"github.com/NzKSO/goExchange/exchange/stex/localdata"
	"github.com/NzKSO/goExchange/exchange/stex/model"
	socketio "github.com/NzKSO/socketio-client-go"
	"github.com/NzKSO/socketio-client-go/protocol"
	"github.com/gorilla/websocket"
)

const (
	restfulEndpoint  = "https://api3.stex.com"
	socketioEndpoint = "wss://socket.stex.com:443"

	allCurrencyPairsURI = "/public/currency_pairs/list/ALL"
)

// Stex represents object stex
type Stex struct {
	allCurrencyPairs map[string]*model.CurrencyPair
	socket           *socketio.SocketClient
	proxy            func(*http.Request) (*url.URL, error)
}

// NewStex returns an instance of exchange stex
func NewStex() exchange.Exchange {
	return &Stex{
		allCurrencyPairs: localdata.AllCurrencyPairs,
	}
}

// Subscribe implements subscribing data from exchange stex
func (s *Stex) Subscribe(sub exchange.Subscriber, symbols ...string) <-chan interface{} {
	ch := make(chan interface{})

	go func() {
		defer func() {
			recover()
		}()

		if s.socket == nil {
			var err error
			s.socket, err = socketio.Socket(socketioEndpoint, &protocol.WebSocketTransport{
				Dialer: websocket.Dialer{
					Proxy: s.proxy,
				},
			})
			if err != nil {
				log.Println(err)
				close(ch)
				return
			}

			s.socket.On(socketio.EventConnect, func(args ...interface{}) {
				log.Println("Connected!")
			})

			s.socket.Connect(nil)
		}

		ctx := context.WithValue(context.Background(), exchange.ContextKey("socket"), s.socket)

		var invalid int
		ids := make([]string, len(symbols))

		for _, symbol := range symbols {
			currencyPair, ok := s.allCurrencyPairs[s.ConvertCurrencyPair(symbol)]
			if !ok {
				log.Printf("%s doesn't exist!!", symbol)
				invalid++
				continue
			}

			ids = append(ids, strconv.FormatInt(int64(currencyPair.ID), 10))
		}

		if invalid == len(symbols) {
			close(ch)
			return
		}

		sub.Subscribe(ctx, ids, ch)
	}()

	return ch
}

// ConvertCurrencyPair converts currency pair into applicable symbol for stex
func (s *Stex) ConvertCurrencyPair(currencyPair string) string {
	return strings.Replace(strings.ToUpper(currencyPair), "/", "_", 1)
}

// SetProxy sets proxy
func (s *Stex) SetProxy(proxy func(*http.Request) (*url.URL, error)) {
	s.proxy = proxy
}
