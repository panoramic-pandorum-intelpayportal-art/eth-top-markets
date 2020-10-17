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

	info := []string{"Name", "symbol", "address", "owner", "decimals", "coingecko", "logo"}

	writer := csv.NewWriter(file)
	defer writer.Flush()
	//csvWriter.Comma = ','
	err = writer.Write(info)

	for _, token := range market.Tokens {
		if token.Symbol == "ETH" || token.Symbol == "USDT" || token.Symbol == "BVOL" || token.Symbol == "IBVOL" || token.Symbol == "AMPL" || token.Symbol == "XAMP" {
			continue
		}

		// do not list tokens which are not on coingecko or don't have an address
		if token.Coingecko == "" || token.Address == "" {
			continue
		}

		tokenInfo := []string{token.Name, token.Symbol, "https://etherscan.io/address/"+token.Address, token.Owner, token.Decimals, "https://www.coingecko.com/en/coins/" + token.Coingecko, "https://ethplorer.io" + token.Image}
		err = writer.Write(tokenInfo)
	}
}
