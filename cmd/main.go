package main

import (
	"fmt"
	"log"
	"os"

	av "github.com/cmckee-dev/go-alpha-vantage"
)

func getQuote(symbol string) (float64, error) {
	client := av.NewClient(os.Getenv("AV_API_KEY"))
	values, err := client.StockTimeSeries(av.TimeSeriesDaily, symbol)
	if err != nil {
		return 0.0, err
	}
	return values[len(values)-1].Close, nil
}

func main() {
	q, err := getQuote("GOOG")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(q)
}
