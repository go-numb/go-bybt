package rest

import (
	"fmt"
	"net/url"

	jsoniter "github.com/json-iterator/go"
	"github.com/valyala/fasthttp"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Response struct {
	Code    string      `json:"code,omitempty"`
	Msg     string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Success bool        `json:"success"`
}

func (p *Client) request(req Requester, results interface{}) error {
	res, err := p.do(req)
	if err != nil {
		return err
	}

	if err := decode(res, results); err != nil {
		return err
	}
	return nil
}

func (p *Client) newRequest(r Requester) *fasthttp.Request {
	// avoid Pointer's butting
	u, _ := url.ParseRequestURI(ENDPOINT)
	u.Path = u.Path + r.Path()
	u.RawQuery = r.Query()

	req := fasthttp.AcquireRequest()
	req.Header.SetMethod(r.Method())
	req.SetRequestURI(u.String())
	body := r.Payload()
	req.SetBody(body)

	if p.Auth != nil {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("bybtSecret", p.Auth.Key)
	}

	return req
}

func (c *Client) do(r Requester) (*fasthttp.Response, error) {
	req := c.newRequest(r)

	// fasthttp for http2.0
	res := fasthttp.AcquireResponse()
	err := c.HTTPC.DoTimeout(req, res, c.HTTPTimeout)
	if err != nil {
		return nil, err
	}

	// fmt.Printf("%+v\n", string(res.Body()))
	// no usefull headers

	// fmt.Printf("+%v", res)

	if res.StatusCode() != 200 {
		var r = new(Response)
		if err := json.Unmarshal(res.Body(), r); err != nil {
			return nil, &APIError{
				Status:  res.StatusCode(),
				Message: err.Error(),
			}
		}

		if !r.Success {
			return nil, &APIError{
				Status:  res.StatusCode(),
				Message: r.Msg,
			}
		}
	}
	return res, nil
}

func decode(res *fasthttp.Response, out interface{}) error {
	var r = new(Response)
	r.Data = out

	if err := json.Unmarshal(res.Body(), r); err != nil {
		return err
	}
	if !r.Success {
		return fmt.Errorf("decode error")
	}
	return nil
}
