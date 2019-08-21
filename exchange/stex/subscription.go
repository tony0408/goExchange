package stex

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/NzKSO/goExchange/exchange/stex/model"
	socketio "github.com/NzKSO/socketio-client-go"
	"github.com/tony0408/goExchange/exchange"
)

const (
	sellChannelPrefix = "sell_data"
	buyChannelPrefix  = "buy_data"
	subscribeEvent    = "subscribe"
	orderBookEvent    = "App\\Events\\GlassRowChanged"

	socketKey = exchange.ContextKey("socket")
)

func emit(client *socketio.SocketClient, event string, args ...interface{}) {
	for _, v := range args {
		client.Emit(event, v)
	}
}

// SubscribeOrderBook subscribes order book identified by identifiers from stex
func SubscribeOrderBook(ctx context.Context, identifiers interface{}, out chan<- interface{}) {
	log.Printf("ctx: %v, iden: %v, out: %v", ctx, identifiers, out)
	ids, ok := identifiers.([]string)
	if !ok {
		panic(fmt.Sprintf("unexpected type %T for identifiers, expect %T", identifiers, ids))
	}

	socket, _ := ctx.Value(socketKey).(*socketio.SocketClient)
	if socket == nil {
		panic("socket is nil")
	}
	log.Println("1")
	orderBookEventHandler := func(args ...interface{}) {
		log.Println("2")
		bytes, err := json.Marshal(args[1])
		if err != nil {
			log.Println(err)
			return
		}

		var msg model.OrderBook
		if err = json.Unmarshal(bytes, &msg); err != nil {
			log.Println(err)
			return
		}

		/* if v, _ := strconv.ParseFloat(msg.Amount, 64); v == 0 {
			fmt.Printf("%+v, Skipped\n", msg)
			return
		} */
		log.Println("3")
		channelName, ok := args[0].(string)
		if ok && strings.HasPrefix(channelName, sellChannelPrefix) {
			msg.Amount = "-" + msg.Amount
		}
		log.Printf("out<-&msg")
		out <- &msg
	}

	socket.On(orderBookEvent, orderBookEventHandler)

	for _, id := range ids {
		emit(socket, subscribeEvent, map[string]interface{}{
			"channel": buyChannelPrefix + id,
			"auth":    nil,
		}, map[string]interface{}{
			"channel": sellChannelPrefix + id,
			"auth":    nil,
		})
	}
}
