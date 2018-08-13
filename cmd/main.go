package main

import (
	"fmt"
	"log"
	"os"

	av "github.com/cmckee-dev/go-alpha-vantage"
)

// Environment vars
var avAPIKey = os.Getenv("AV_API_KEY")
var avSymbol = os.Getenv("AV_SYMBOL")

// getLatestQuote fetches the latest closing quote for a symbol.
func getLatestQuote(client *av.Client, symbol string) (float64, error) {
	values, err := client.StockTimeSeries(av.TimeSeriesDaily, symbol)
	if err != nil {
		return 0.0, err
	}
	return values[len(values)-1].Close, nil
}

// main.
func main() {
	client := av.NewClient(avAPIKey)
	q, err := getLatestQuote(client, avSymbol)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(q)
}
