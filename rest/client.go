package rest

import (
	"time"

	"github.com/go-numb/go-bybt/auth"

	"github.com/valyala/fasthttp"
)

const ENDPOINT = "http://open-api.bybt.com/api/pro/v1"

type Client struct {
	Auth *auth.Config

	HTTPC       *fasthttp.Client
	HTTPTimeout time.Duration
}

func New(auth *auth.Config) *Client {
	hc := new(fasthttp.Client)

	return &Client{
		Auth:        auth,
		HTTPC:       hc,
		HTTPTimeout: 5 * time.Second,
	}
}
