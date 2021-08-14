package liquidation

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type RequestWithChart struct {
	Symbol   string `url:"symbol"`
	TimeType string `url:"timeType"`
}

type ResponseWithChart []WithChart

type WithChart struct {
	BuyVolUsd          float64 `json:"buyVolUsd"`
	SellVolUsd         float64 `json:"sellVolUsd"`
	BuyQty             float64 `json:"buyQty"`
	SellQty            float64 `json:"sellQty"`
	BuyTurnoverNumber  int     `json:"buyTurnoverNumber"`
	SellTurnoverNumber int     `json:"sellTurnoverNumber"`
	TurnoverNumber     int     `json:"turnoverNumber"`
	Price              float64 `json:"price"`
	List               []struct {
		ExchangeName       string  `json:"exchangeName"`
		BuyVolUsd          float64 `json:"buyVolUsd"`
		SellVolUsd         float64 `json:"sellVolUsd"`
		TurnoverNumber     int     `json:"turnoverNumber"`
		BuyTurnoverNumber  int     `json:"buyTurnoverNumber"`
		SellTurnoverNumber int     `json:"sellTurnoverNumber"`
		SellQty            float64 `json:"sellQty"`
		BuyQty             float64 `json:"buyQty"`
	} `json:"list"`
	CreateTime int64 `json:"createTime"`
}

func (req *RequestWithChart) Path() string {
	return "/futures/liquidation/detail/chart"
}

func (req *RequestWithChart) Method() string {
	return http.MethodGet
}

func (req *RequestWithChart) Query() string {
	value, _ := query.Values(req)
	return value.Encode()
}

func (req *RequestWithChart) Payload() []byte {
	return nil
}
