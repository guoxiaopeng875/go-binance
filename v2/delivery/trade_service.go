package delivery

import (
	"context"
	"encoding/json"
	"net/http"
)

// ListAccountTradeService define account trade list service
type ListAccountTradeService struct {
	c         *Client
	symbol    string
	startTime *int64
	endTime   *int64
	fromID    *int64
	limit     *int
}

// Symbol set symbol
func (s *ListAccountTradeService) Symbol(symbol string) *ListAccountTradeService {
	s.symbol = symbol
	return s
}

// StartTime set startTime
func (s *ListAccountTradeService) StartTime(startTime int64) *ListAccountTradeService {
	s.startTime = &startTime
	return s
}

// EndTime set endTime
func (s *ListAccountTradeService) EndTime(endTime int64) *ListAccountTradeService {
	s.endTime = &endTime
	return s
}

// FromID set fromID
func (s *ListAccountTradeService) FromID(fromID int64) *ListAccountTradeService {
	s.fromID = &fromID
	return s
}

// Limit set limit
func (s *ListAccountTradeService) Limit(limit int) *ListAccountTradeService {
	s.limit = &limit
	return s
}

// Do send request
func (s *ListAccountTradeService) Do(ctx context.Context, opts ...RequestOption) (res []*AccountTrade, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/dapi/v1/userTrades",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)
	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}
	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}
	if s.fromID != nil {
		r.setParam("fromId", *s.fromID)
	}
	if s.limit != nil {
		r.setParam("limit", *s.limit)
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return []*AccountTrade{}, err
	}
	res = make([]*AccountTrade, 0)
	err = json.Unmarshal(data, &res)
	if err != nil {
		return []*AccountTrade{}, err
	}
	return res, nil
}

// AccountTrade define account trade
type AccountTrade struct {
	Symbol          string           `json:"symbol"`          // 交易对
	Id              int              `json:"id"`              // 交易ID
	OrderId         int              `json:"orderId"`         // 订单ID
	Pair            string           `json:"pair"`            // 标的交易对
	Side            SideType         `json:"side"`            // 买卖方向
	Price           string           `json:"price"`           // 成交价
	Qty             string           `json:"qty"`             // 成交量(张数)
	RealizedPnl     string           `json:"realizedPnl"`     // 实现盈亏
	MarginAsset     string           `json:"marginAsset"`     // 保证金币种
	BaseQty         string           `json:"baseQty"`         // 成交额(标的数量)
	Commission      string           `json:"commission"`      // 手续费
	CommissionAsset string           `json:"commissionAsset"` // 手续费单位
	Time            int64            `json:"time"`            // 时间
	PositionSide    PositionSideType `json:"positionSide"`    // 持仓方向
	Buyer           bool             `json:"buyer"`           // 是否是买方
	Maker           bool             `json:"maker"`           // 是否是挂单方
}
