package binance

import (
	"context"
	"net/http"
)

// GetSimpleEarnAccountService get simple earn account info
type GetSimpleEarnAccountService struct {
	c *Client
}

// Do send request
func (s *GetSimpleEarnAccountService) Do(ctx context.Context, opts ...RequestOption) (res *SimpleEarnAccount, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/simple-earn/account",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opts...)
	if err != nil {
		return nil, err
	}
	res = new(SimpleEarnAccount)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// SimpleEarnAccount define simple-earn account info
type SimpleEarnAccount struct {
	TotalAmountInBTC          string `json:"totalAmountInBTC"`
	TotalAmountInUSDT         string `json:"totalAmountInUSDT"`
	TotalFlexibleAmountInBTC  string `json:"totalFlexibleAmountInBTC"`
	TotalFlexibleAmountInUSDT string `json:"totalFlexibleAmountInUSDT"`
	TotalLockedInBTC          string `json:"totalLockedInBTC"`
	TotalLockedInUSDT         string `json:"totalLockedInUSDT"`
}
