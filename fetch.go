package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func FetchConversion() ExchangeResponse {
	url := fmt.Sprintf("https://openexchangerates.org/api/latest.json?app_id=%s", DefaultConfig.EXCHANGE_RATE_APP_ID)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("fetch: error making http request: %s\n ", err)
	}

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("fetch: could not read response body: %s\n", err)
	}

	var exchangeResponse ExchangeResponse
	err = json.Unmarshal(resBody, &exchangeResponse)
	if err != nil {
		log.Fatalf("fetch: could not unmarshal response body: %s\n", err)
	}

	return exchangeResponse
}
