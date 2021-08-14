package rest_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/go-numb/go-bybt/auth"
	"github.com/go-numb/go-bybt/rest"
	"github.com/go-numb/go-bybt/rest/futures/fr"
	"github.com/go-numb/go-bybt/rest/futures/liquidation"
	"github.com/go-numb/go-bybt/rest/futures/oi"
	"github.com/stretchr/testify/assert"
)

func TestOpenInterest(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("BYBTAPIKEY")))
	res, err := c.OpenInterest(&oi.Request{
		Symbol:   "ETH",
		Interval: 0,
	})
	assert.NoError(t, err)

	for i := 0; i < len(*res); i++ {
		fmt.Printf("%s,	%.5f,	%.5f\n", (*res)[i].ExchangeName, (*res)[i].H1OIChangePercent, (*res)[i].H4OIChangePercent)
	}
}

func TestOpenInterestWithChart(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("BYBTAPIKEY")))
	res, err := c.OpenInterestWithChart(&oi.RequestWithChart{
		Symbol:   "ETH",
		Interval: 1,
	})
	assert.NoError(t, err)

	fmt.Printf("%+v", res)
}

func TestLiquidation(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("BYBTAPIKEY")))
	res, err := c.Liquidation(&liquidation.Request{
		ExName: "Binance",
		Symbol: "ETH",
	})
	assert.NoError(t, err)

	fmt.Printf("%+v", res)
}

func TestLiquidationWithChart(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("BYBTAPIKEY")))
	res, err := c.LiquidationWithChart(&liquidation.RequestWithChart{
		TimeType: "9",
		Symbol:   "ETH",
	})
	assert.NoError(t, err)

	fmt.Printf("%+v", res)
}

func TestFundingRateWithChart(t *testing.T) {
	c := rest.New(auth.New(os.Getenv("BYBTAPIKEY")))
	res, err := c.FundingRateWithChart(&fr.RequestWithChart{
		Symbol: "ETH",
	})
	assert.NoError(t, err)

	fmt.Printf("%+v", res)
}
