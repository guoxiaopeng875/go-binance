package delivery

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type tradeServiceTestSuite struct {
	baseTestSuite
}

func TestTradeService(t *testing.T) {
	suite.Run(t, new(tradeServiceTestSuite))
}

func (s *tradeServiceTestSuite) TestAccountTradeList() {
	data := []byte(`[
		{
		  "symbol": "BTCUSD_200626",
		  "id": 6,
		  "orderId": 28,
		  "pair": "BTCUSD",
		  "side": "SELL",
		  "price": "8800",
		  "qty": "1",
		  "realizedPnl": "0",
		  "marginAsset": "BTC",
		  "baseQty": "0.01136364",
		  "commission": "0.00000454",
		  "commissionAsset": "BTC",
		  "time": 1590743483586,
		  "positionSide": "BOTH",
		  "buyer": false,
		  "maker": false
	}
	]`)
	s.mockDo(data, nil)
	defer s.assertDo()

	symbol := "BTCUSD_200626"
	startTime := int64(1569514978020)
	endTime := int64(1569514978021)
	fromID := int64(698759)
	limit := 3
	s.assertReq(func(r *request) {
		e := newSignedRequest().setParams(params{
			"symbol":    symbol,
			"startTime": startTime,
			"endTime":   endTime,
			"fromId":    fromID,
			"limit":     limit,
		})
		s.assertRequestEqual(e, r)
	})

	trades, err := s.client.NewListAccountTradeService().Symbol(symbol).
		StartTime(startTime).EndTime(endTime).FromID(fromID).Limit(limit).Do(newContext())
	r := s.r()
	r.NoError(err)
	r.Len(trades, 1)
	e := &AccountTrade{
		Symbol:          symbol,
		Id:              6,
		OrderId:         28,
		Pair:            "BTCUSD",
		Side:            SideTypeSell,
		Price:           "8800",
		Qty:             "1",
		RealizedPnl:     "0",
		MarginAsset:     "BTC",
		BaseQty:         "0.01136364",
		Commission:      "0.00000454",
		CommissionAsset: "BTC",
		Time:            1590743483586,
		PositionSide:    PositionSideTypeBoth,
		Buyer:           false,
		Maker:           false,
	}
	s.assertAccountTradeEqual(e, trades[0])
}

func (s *tradeServiceTestSuite) assertAccountTradeEqual(e, a *AccountTrade) {
	r := s.r()
	r.Equal(e.Id, a.Id, "Id")
	r.Equal(e.Pair, a.Pair, "Pair")
	r.Equal(e.Buyer, a.Buyer, "Buyer")
	r.Equal(e.Commission, a.Commission, "Commission")
	r.Equal(e.CommissionAsset, a.CommissionAsset, "CommissionAsset")
	r.Equal(e.MarginAsset, a.MarginAsset, "MarginAsset")
	r.Equal(e.Maker, a.Maker, "Maker")
	r.Equal(e.OrderId, a.OrderId, "OrderId")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Qty, a.Qty, "Qty")
	r.Equal(e.BaseQty, a.BaseQty, "BaseQty")
	r.Equal(e.RealizedPnl, a.RealizedPnl, "RealizedPnl")
	r.Equal(e.Side, a.Side, "Side")
	r.Equal(e.PositionSide, a.PositionSide, "PositionSide")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.Time, a.Time, "Time")
}
