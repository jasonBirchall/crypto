package api

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Query(c string) ([]byte, error) {
	apiUrl := "https://api.coinranking.com/v1/public/coins?base=gbp&prefix=" + c
	client := http.Client{
		Timeout: time.Second * 2,
	}

	req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "crypto-tracker")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}
