package main

import (
	"fmt"
	"log"
	"strings"

	gecko "github.com/superoo7/go-gecko/v3"
)

type GeckoApi struct {
	URL string
}

func NewGeckoApi() *GeckoApi {
	return &GeckoApi{
		URL: "https://api.coingecko.com/api/v3",
	}
}

var cg = gecko.NewClient(nil)
var ticker = map[string]string{
	"BTC":  "bitcoin",
	"ETH":  "ethereum",
	"USDT": "usd",
}

func (g *GeckoApi) LoadCoins() (int, error) {
	list, err := cg.CoinsList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Available coins:")

	return len(*list), nil
}

func (g *GeckoApi) GetRates(from, to string) (float64, error) {
	fromParam := []string{strings.ToLower(from)}
	toParam := []string{strings.ToLower(to)}

	sp, err := cg.SimplePrice(fromParam, toParam)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch rates: %w", err)
	}

	if rate, exists := (*sp)[from][to]; exists {
		fmt.Println(from, rate, to)
		return float64(rate), nil
	}

	return 0, fmt.Errorf("conversion rate not found")
}
