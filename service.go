package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GetCurrentPrice fetches the current price of a crypto asset from CoinGecko
func GetCurrentPrice(cryptoName string) (float64, error) {
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", cryptoName)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var result map[string]map[string]float64
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}

	if price, ok := result[cryptoName]; ok {
		return price["usd"], nil
	}

	return 0, fmt.Errorf("crypto asset not found")
}
