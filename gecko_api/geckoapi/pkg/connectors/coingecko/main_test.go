package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadCoins(t *testing.T) {
	api := NewGeckoApi()

	// AMOUNT OF COINS
	count, err := api.LoadCoins()

	// CHECKING ERROR
	assert.NoError(t, err)

	// CHECKING IF AMOUNT > 0
	assert.True(t, count > 0, "Expected to have available coins")
}

func TestGetRates(t *testing.T) {
	api := NewGeckoApi()

	// GETTING RATE FOR BTC
	price, err := api.GetRates(ticker["BTC"], ticker["USDT"])

	assert.NoError(t, err)

	assert.True(t, price > 0, "Price should be greater than zero")

	// GETTING RATE FOR ETH
	price, err = api.GetRates(ticker["ETH"], ticker["USDT"])

	assert.NoError(t, err)

	assert.True(t, price > 0, "Price should be greater than zero")
}

// TEST FOR INCORRECT CURR
func TestGetRatesError(t *testing.T) {
	api := NewGeckoApi()

	// RANDOM RATE
	_, err := api.GetRates("nonexistentcoin", "USDT")

	// EXPECTED ERROR
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "conversion rate not found")
}
