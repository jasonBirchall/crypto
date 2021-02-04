// Package api allows users to query a popular coin ranking API's
// which is used in various commands.
package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// Query takes a coin shorthand value as a string and queries an endpoint.
// It outputs a slice of bytes and an error value if there is one.
func Query(c string) ([]byte, error) {
	apiUrl := "https://api.coinranking.com/v1/public/coins?base=gbp&prefix=" + c
	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		log.Panic(err)
	}

	req.Header.Set("User-Agent", "crypto-tracker")

	res, err := client.Do(req)
	if err != nil {
		log.Panic("Unable to contact coinranking API")
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Panic(err)
	}

	return body, nil
}
