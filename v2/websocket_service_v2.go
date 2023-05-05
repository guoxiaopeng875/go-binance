package binance

import (
	"fmt"
	"github.com/adshao/go-binance/v2/common"
	"os"
	"time"
)

// WsUserDataServeV2 serve user data handler with listen key
func WsUserDataServeV2(listenKey string, handler WsUserDataHandler) *common.WsConn {
	endpoint := fmt.Sprintf("%s/%s", getWsEndpoint(), listenKey)
	wsHandler := func(message []byte) error {
		j, err := newJSON(message)
		if err != nil {
			return err
		}

		event := new(WsUserDataEvent)

		err = json.Unmarshal(message, event)
		if err != nil {
			return err
		}

		switch UserDataEventType(j.Get("e").MustString()) {
		case UserDataEventTypeOutboundAccountPosition:
			err = json.Unmarshal(message, &event.AccountUpdate)
			if err != nil {
				return err
			}
		case UserDataEventTypeBalanceUpdate:
			err = json.Unmarshal(message, &event.BalanceUpdate)
			if err != nil {
				return err
			}
		case UserDataEventTypeExecutionReport:
			err = json.Unmarshal(message, &event.OrderUpdate)
			if err != nil {
				return err
			}
			// Unmarshal has case sensitive problem
			event.TransactionTime = j.Get("T").MustInt64()
			event.OrderUpdate.TransactionTime = j.Get("T").MustInt64()
			event.OrderUpdate.Id = j.Get("i").MustInt64()
			event.OrderUpdate.TradeId = j.Get("t").MustInt64()
			event.OrderUpdate.FeeAsset = j.Get("N").MustString()
		case UserDataEventTypeListStatus:
			err = json.Unmarshal(message, &event.OCOUpdate)
			if err != nil {
				return err
			}
		}

		handler(event)

		return nil
	}
	return common.NewWsBuilder().
		ProxyUrl(os.Getenv("HTTPS_PROXY")).
		ProtoHandleFunc(wsHandler).
		Heartbeat(func() []byte {
			return []byte("1")
		}, time.Second*30).
		AutoReconnect().
		WsUrl(endpoint).
		Build()
}
