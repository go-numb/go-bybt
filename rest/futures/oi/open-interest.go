package oi

import (
	"net/http"

	"github.com/google/go-querystring/query"
)

type Request struct {
	Symbol   string `url:"symbol"`
	Interval int    `url:"interval"`
}

type Response []OI

type OI struct {
	ExchangeName       string  `json:"exchangeName"`
	Symbol             string  `json:"symbol"`
	Price              float64 `json:"price"`
	OpenInterest       float64 `json:"openInterest"`
	OpenInterestAmount float64 `json:"openInterestAmount"`
	VolUsd             float64 `json:"volUsd"`
	Rate               float64 `json:"rate"`
	H24Change          float64 `json:"h24Change"`
	H1OIChangePercent  float64 `json:"h1OIChangePercent"`
	H4OIChangePercent  float64 `json:"h4OIChangePercent"`
}

func (req *Request) Path() string {
	return "/futures/openInterest"
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
