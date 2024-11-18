package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadCoins(t *testing.T) {
	// AMOUNT OF COINS
	count, err := LoadCoins()

	// CHECKING ERROR
	assert.NoError(t, err)

	// CHECKING IF AMOUNT > 0
	assert.True(t, count > 0, "Expected to have available coins")
}

func TestGetRates(t *testing.T) {
	// GETTING RATE FOR BTC
	price, err := GetRates(ticker["BTC"], ticker["USDT"])

	assert.NoError(t, err)

	assert.True(t, price > 0, "Price should be greater than zero")

	// GETTING RATE FOR ETH
	price, err = GetRates(ticker["ETH"], ticker["USDT"])

	assert.NoError(t, err)

	assert.True(t, price > 0, "Price should be greater than zero")
}

// TEST FOR INCORRECT CURR
func TestGetRatesError(t *testing.T) {

	// RANDOM RATE
	_, err := GetRates("nonexistentcoin", "USDT")

	// EXPECTED ERROR
	assert.Error(t, err)
	assert.Equal(t, err.Error(), "conversion rate not found")
}
