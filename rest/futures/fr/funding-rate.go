package fr

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type RequestWithChart struct {
	Type   string `url:"type,omitempty"`
	Symbol string `url:"symbol"`
}

type ResponseWithChart Fundings

type Fundings struct {
	PriceList []float64            `json:"priceList"`
	DataMap   map[string][]float64 `json:"dataMap"`
	FrDataMap map[string][]float64 `json:"frDataMap"`
	DateList  []int64              `json:"dateList"`
}

func (req *RequestWithChart) Path() string {
	return "/futures/funding_rates_chart"
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
