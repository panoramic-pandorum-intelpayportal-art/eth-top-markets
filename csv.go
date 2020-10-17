package main

import (
	"encoding/csv"
	"os"
)

func csvWriteEthTopMarkets(market *EthMarket) {
	os.Remove("tokens.csv")
	file, err := os.OpenFile("tokens.csv", os.O_CREATE|os.O_WRONLY, 0777)
	defer file.Close()

	if err != nil {
		os.Exit(1)
	}

	info := []string{"Name", "symbol", "contract address", "decimals", "coingecko", "logo"}

	writer := csv.NewWriter(file)
	defer writer.Flush()
	//csvWriter.Comma = ','
	err = writer.Write(info)

	for _, token := range market.Tokens {
		tokenInfo := []string{token.Name, token.Symbol, "https://etherscan.io/address/"+token.Address, token.Decimals, "https://www.coingecko.com/en/coins/" + token.Coingecko, "https://ethplorer.io" + token.Image}

		if token.Symbol == "ETH" || token.Symbol == "USDT" {
			continue
		}
		err = writer.Write(tokenInfo)
	}
}
