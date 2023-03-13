package futures

import (
	"encoding/json"
	"fmt"
	"github.com/adshao/go-binance/v2/common"
	"os"
	"strings"
	"time"
)

// WsUserDataServeV2 serve user data handler with listen key
func WsUserDataServeV2(listenKey string, handler WsUserDataHandler) *common.WsConn {
	endpoint := fmt.Sprintf("%s/%s", getWsEndpoint(), listenKey)
	wsHandler := func(message []byte) error {
		event := new(WsUserDataEvent)
		if err := json.Unmarshal(message, event); err != nil {
			return err
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

// WsMiniMarketTickerServeV2 serve websocket that pushes 24hr rolling window mini-ticker statistics for a single symbol.
func WsMiniMarketTickerServeV2(symbol string, handler WsMiniMarketTickerHandler) *common.WsConn {
	endpoint := fmt.Sprintf("%s/%s@miniTicker", getWsEndpoint(), strings.ToLower(symbol))
	wsHandler := func(message []byte) error {
		event := new(WsMiniMarketTickerEvent)
		if err := json.Unmarshal(message, event); err != nil {
			return err
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
