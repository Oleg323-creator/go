package main

import (
	"fmt"
	"geckoapi1/pkg/connectors/coingecko"
	"geckoapi1/pkg/connectors/crypto_compare"
)

type ConnectorAPI interface {
	LoadCoins() (int, error)
	GetRates(from, to string) (map[string]interface{}, error)
}

func releaseInter(c ConnectorAPI, from string, to string) {
	fmt.Println(c.LoadCoins())
	fmt.Println(c.GetRates(from, to))
}

type Connector struct {
	ConnectorAPI
}

func NewConnector(conType string) (*Connector, error) {
	const coingeckoType = "Coingecko"
	const cryptoCompType = "Crypto Compare"

	if conType == coingeckoType {
		return &Connector{
			coingecko.NewGeckoApi(),
		}, nil
	} else if conType == cryptoCompType {
		return &Connector{
			crypto_compare.NewCryptoCompareAPI(),
		}, nil
	} else {
		return nil, fmt.Errorf("unknown connector type")
	}
}

func main() {

	//GECKO IMPLEMENTATION
	conn, err := NewConnector("Coingecko")
	if err != nil {
		return
	}
	_, err = conn.LoadCoins()
	if err != nil {
		return
	}

	rate, err := conn.GetRates("BTC", "USDT")
	if err != nil {
		return
	}
	fmt.Println(rate)

	//CRYPTO_COMPARE IMPLEMENTATION
	conn, err = NewConnector("Crypto Compare")
	if err != nil {
		return
	}
	_, err = conn.LoadCoins()
	if err != nil {
		return
	}

	rate, err = conn.GetRates("BTC", "USDT")
	if err != nil {
		return
	}
	fmt.Println(rate)
}
