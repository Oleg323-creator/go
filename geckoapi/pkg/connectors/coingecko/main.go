package main

import (
	"fmt"
	"log"
	"strings"

	gecko "github.com/superoo7/go-gecko/v3"
)

var cg = gecko.NewClient(nil)
var ticker = map[string]string{
	"BTC":  "bitcoin",
	"ETH":  "ethereum",
	"USDT": "usd",
}

func LoadCoins() (int, error) {
	list, err := cg.CoinsList()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Available coins:")

	return len(*list), nil
}

func GetRates(from, to string) (float64, error) {
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

func main() {
	load, err := LoadCoins()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(load)

	get, err := GetRates(ticker["BTC"], ticker["USDT"])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(get)
}
