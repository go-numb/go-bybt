package liquidation

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Request struct {
	Symbol string `url:"symbol,omitempty"`
	ExName string `url:"exName,omitempty"`
}

type Response Liquidation

type Liquidation struct {
	SellQty float64 `json:"sellQty"`
	BuyQty  float64 `json:"buyQty"`

	PriceList []float64 `json:"priceList"`
	SellList  []float64 `json:"sellList"`
	BuyList   []float64 `json:"buyList"`
	VolList   []float64 `json:"volList"`
	NumList   []int     `json:"numList"`
	DateList  []int64   `json:"dateList"`
}

func (req *Request) Path() string {
	return "/futures/liquidation_chart"
}

func (req *Request) Method() string {
	return http.MethodGet
}

func (req *Request) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *Request) Payload() []byte {
	return nil
}
