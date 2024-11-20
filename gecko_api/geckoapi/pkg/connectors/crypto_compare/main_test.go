package main

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestGetRatesFromCC(t *testing.T) {
	api := NewCryptoCompareAPI()

	// GETTING RATE FOR BTC
	paramsBTC := map[string]string{
		"fsym=":  "BTC",
		"tsyms=": "USD",
	}
	data, err := api.GetRatesFromCC("/price?", paramsBTC)
	log.Println(data)

	assert.NoError(t, err, "Unexpected error occurred")

	if assert.NotNil(t, data, "Expected data but got nil") {
		assert.Contains(t, data, "USD", "Expected USD in response")
	}

	// GETTING RATE FOR ETH
	paramsETH := map[string]string{
		"fsym=":  "ETH",
		"tsyms=": "USD",
	}
	data, err = api.GetRatesFromCC("/price?", paramsETH)
	log.Println(data)

	assert.NoError(t, err, "Unexpected error occurred")

	if assert.NotNil(t, data, "Expected data but got nil") {
		assert.Contains(t, data, "USD", "Expected USD in response")
	}

}

// TEST FOR INCORRECT CURR
func TestGetRatesError(t *testing.T) {
	api := NewCryptoCompareAPI()

	// RANDOM RATE
	paramsRandom := map[string]string{
		"fsym=":  "nonexistentcoin",
		"tsyms=": "USD",
	}
	data, err := api.GetRatesFromCC("/price?", paramsRandom)
	log.Println(data)

	// EXPECTED ERROR
	if assert.Error(t, err, "Expected an error but got nil") {
		assert.Contains(t, err.Error(), "market does not exist for this coin pair", "Error message mismatch")
	}

	assert.Nil(t, data, "Expected data to be nil on error")
}
