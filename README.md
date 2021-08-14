# go-bybt
Bybt of General information for Crypto.
This client is wrapper of Bybt.

## Description
go-bybt is a go client follow [bybt API Document](https://bybt.gitbook.io/bybt/api-key).

**Supported**
- [x] Gets Open Interest/With Chart
- [x] Gets Liquidation/With Chart
- [x] Get Funding Rate




## Usage
``` sample.go
package main

import (
	"fmt"

	"github.com/go-numb/go-bybt/auth"
	"github.com/go-numb/go-bybt/rest"
	"github.com/go-numb/go-bybt/rest/futures/fr"
	"github.com/go-numb/go-bybt/rest/futures/liquidation"
	"github.com/go-numb/go-bybt/rest/futures/oi"
)


func main() {
    // Can be obtained by registering with bybt
    key := "bybt apikey"
	c := rest.New(auth.New(auth.New(key)))
	res, err := c.FundingRateWithChart(&fr.RequestWithChart{
		Symbol: "ETH",
	})
	if err != nil {
        log.Fatal(err)
    }

	fmt.Printf("%+v", res)
}


```

## Author

[@_numbP](https://twitter.com/_numbP)

## License

[MIT](https://github.com/go-numb/go-bybt/blob/master/LICENSE)