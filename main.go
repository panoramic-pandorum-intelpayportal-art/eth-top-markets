package main

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func main() {
	ethMarkets := GetEthereumTokens()
	csvWriteEthTopMarkets(ethMarkets)
}

type EthMarket struct {
	Tokens []struct {
		Address         string  `json:"address"`
		Name            string  `json:"name"`
		Symbol          string  `json:"symbol"`
		Volume          float64 `json:"volume"`
		Cap             float64 `json:"cap"`
		AvailableSupply float64 `json:"availableSupply"`
		Price           struct {
			Rate            float64 `json:"rate"`
			Diff            float64 `json:"diff"`
			Diff7D          float64 `json:"diff7d"`
			Ts              int     `json:"ts"`
			MarketCapUsd    float64 `json:"marketCapUsd"`
			AvailableSupply float64 `json:"availableSupply"`
			Volume24H       float64 `json:"volume24h"`
			Diff30D         float64 `json:"diff30d"`
		} `json:"price"`
		Volume1DCurrent   float64 `json:"volume-1d-current"`
		Volume1DPrevious  float64 `json:"volume-1d-previous"`
		Cap1DCurrent      float64 `json:"cap-1d-current"`
		Cap1DPrevious     float64 `json:"cap-1d-previous"`
		Cap1DPreviousTs   int     `json:"cap-1d-previous-ts"`
		Volume7DCurrent   float64 `json:"volume-7d-current"`
		Volume7DPrevious  float64 `json:"volume-7d-previous"`
		Cap7DCurrent      float64 `json:"cap-7d-current"`
		Cap7DPrevious     float64 `json:"cap-7d-previous"`
		Cap7DPreviousTs   int     `json:"cap-7d-previous-ts"`
		Volume30DCurrent  float64 `json:"volume-30d-current"`
		Volume30DPrevious float64 `json:"volume-30d-previous"`
		Cap30DCurrent     float64 `json:"cap-30d-current"`
		Cap30DPrevious    float64 `json:"cap-30d-previous"`
		Cap30DPreviousTs  int     `json:"cap-30d-previous-ts"`
		Decimals          string  `json:"decimals,omitempty"`
		TotalSupply       string  `json:"totalSupply,omitempty"`
		Owner             string  `json:"owner,omitempty"`
		TxsCount          int     `json:"txsCount,omitempty"`
		TransfersCount    int     `json:"transfersCount,omitempty"`
		LastUpdated       int     `json:"lastUpdated,omitempty"`
		IssuancesCount    int     `json:"issuancesCount,omitempty"`
		HoldersCount      int     `json:"holdersCount,omitempty"`
		Website           string  `json:"website,omitempty"`
		Twitter           string  `json:"twitter,omitempty"`
		Image             string  `json:"image,omitempty"`
		Facebook          string  `json:"facebook,omitempty"`
		Coingecko         string  `json:"coingecko,omitempty"`
		EthTransfersCount int     `json:"ethTransfersCount,omitempty"`
		Reddit            string  `json:"reddit,omitempty"`
		Description       string  `json:"description,omitempty"`
		Telegram          string  `json:"telegram,omitempty"`
	} `json:"tokens"`
	Totals struct {
		Tokens          int     `json:"tokens"`
		TokensWithPrice int     `json:"tokensWithPrice"`
		Cap             float64 `json:"cap"`
		CapPrevious     float64 `json:"capPrevious"`
		Volume24H       float64 `json:"volume24h"`
		VolumePrevious  float64 `json:"volumePrevious"`
		Ts              int     `json:"ts"`
	} `json:"totals"`
}

func GetEthereumTokens() *EthMarket {
	apiEndpoint, err := url.Parse("https://api.ethplorer.io/getTop?apiKey=freekey&criteria=trade&limit=110")
	if err != nil {
		logrus.Fatalf("could not parse url endpoint: %s\n", err)
	}

	client := http.Client{
		Transport:     nil,
		CheckRedirect: nil,
		Jar:           nil,
		Timeout:       time.Second * 15,
	}

	req, err := http.NewRequest(http.MethodGet, apiEndpoint.String(), nil)
	if err != nil {
		logrus.Fatalf("could not create a new request to %s\n", apiEndpoint.Scheme)
	}

	resp, err := client.Do(req)
	if err != nil {
		logrus.Fatalf("client do failed for API endpoint: %s\n", apiEndpoint.String())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Fatalf("Could not read body response from API endpoint %s %s\n", apiEndpoint.String(), err)
	}

	tokens := new(EthMarket)
	err = json.Unmarshal(body, &tokens)

	return tokens
}
