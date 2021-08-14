package oi

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type RequestWithChart struct {
	Symbol   string `url:"symbol"`
	Interval int    `url:"interval"`
}

type ResponseWithChart OIWithChart

type OIWithChart struct {
	DateList  []int            `json:"dateList,omitempty"`
	PriceList []float64        `json:"priceList,omitempty"`
	DataMap   map[string][]int `json:"dataMap,omitempty"`
}

func (req *RequestWithChart) Path() string {
	return "/futures/openInterest/chart"
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
