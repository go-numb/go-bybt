package rest

import (
	"github.com/go-numb/go-bybt/rest/futures/fr"
	"github.com/go-numb/go-bybt/rest/futures/liquidation"
	"github.com/go-numb/go-bybt/rest/futures/oi"
)

func (p *Client) OpenInterest(req *oi.Request) (*oi.Response, error) {
	results := new(oi.Response)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) OpenInterestWithChart(req *oi.RequestWithChart) (*oi.ResponseWithChart, error) {
	results := new(oi.ResponseWithChart)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) Liquidation(req *liquidation.Request) (*liquidation.Response, error) {
	results := new(liquidation.Response)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) LiquidationWithChart(req *liquidation.RequestWithChart) (*liquidation.ResponseWithChart, error) {
	results := new(liquidation.ResponseWithChart)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}

func (p *Client) FundingRateWithChart(req *fr.RequestWithChart) (*fr.ResponseWithChart, error) {
	results := new(fr.ResponseWithChart)
	if err := p.request(req, results); err != nil {
		return nil, err
	}
	return results, nil
}
