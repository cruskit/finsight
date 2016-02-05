package strategy

import (
	"time"
	"github.com/cruskit/finsight/metric"
	"fmt"
	"github.com/cruskit/finsight/analyser"
	"log"
)

type StrategyOutcomes struct {
	PurchaseDates [] time.Time
	SellDates     [] time.Time
	FinalValue    float64
}

func RunMovingAverageCrossover(fastAvgPeriod int, slowAvgPeriod int, dataFile string) *StrategyOutcomes {

	// Get the stock data
	pricesPtr := metric.ReadMetricFromYahooCsv(dataFile)

	fmt.Println(pricesPtr)

	buyDates, sellDates := calculateMovingAverageTxDates(fastAvgPeriod, slowAvgPeriod, pricesPtr)

	fmt.Println("Buy", buyDates)
	fmt.Println("Sell", sellDates)

	outComes := StrategyOutcomes{*buyDates, *sellDates, 0.0}
	calculateStrategyValue(&outComes, pricesPtr)

	return &outComes
}

// calculateMovingAverageTxDates determines the buy and sell dates for a unit based on moving average crossovers
// it returns the buyDates and the sellDates
func calculateMovingAverageTxDates(fastAvgPeriod int, slowAvgPeriod int, pricesPtr *[]metric.Metric) (*[]time.Time, *[]time.Time) {

	fastAvg := analyser.MovingAverage(fastAvgPeriod, pricesPtr)
	slowAvg := analyser.MovingAverage(slowAvgPeriod, pricesPtr)

	// Make the avgs start on the same date
	fastAvg = fastAvg[slowAvgPeriod - fastAvgPeriod:]

	fmt.Println("first fast avg: ", fastAvg[0])
	fmt.Println("first slow avg: ", slowAvg[0])

	// Work out the buy and sell dates
	buy := make([]time.Time, 0)
	sell := make([]time.Time, 0)

	for i := 1; i < len(fastAvg); i++ {

		if fastAvg[i].Value > slowAvg[i].Value && fastAvg[i - 1].Value < slowAvg[i - 1].Value {
			// Low to high crossover - buy
			buy = append(buy, fastAvg[i].Time)

		} else if fastAvg[i].Value < slowAvg[i].Value && fastAvg[i - 1].Value > slowAvg[i - 1].Value && len(buy) > 0 {
			// High to low crossover - sell, can only sell if we've previously bought
			sell = append(sell, fastAvg[i].Time)
		}

	}

	return &buy, &sell
}

func calculateStrategyValue(outcomesPtr *StrategyOutcomes, pricesPtr *[]metric.Metric) {

	var value float64 = 1000.0
	var numShares float64 = 0.0
	inMarket := false


	for i := 0; i < len((*outcomesPtr).PurchaseDates); i++ {

		// Buy the shares
		inMarket = true
		foundBuyPrice := false;

		for j := 0; j < len(*pricesPtr); j++ {

			if ((*outcomesPtr).PurchaseDates[i].Equal((*pricesPtr)[j].Time)) {
				foundBuyPrice = true
				numShares = value / (*pricesPtr)[j].Value
				fmt.Printf("%v Shares units at %v for %v on %v\n",
					numShares, (*pricesPtr)[j].Value, value, (*pricesPtr)[j].Time)
			}
		}
		if !foundBuyPrice {
			log.Panicln("Unable to find price for buy date: ", (*outcomesPtr).PurchaseDates[i])
		}

		// Check if we sold this as well
		if (i < len((*outcomesPtr).SellDates)) {
			inMarket = false
			foundSellPrice := false

			for j := 0; j < len((*pricesPtr)); j++ {
				if ((*outcomesPtr).SellDates[i] == (*pricesPtr)[j].Time) {
					foundSellPrice = true
					value = numShares * (*pricesPtr)[j].Value
					fmt.Printf("%v units sold at %v for %v on %v\n",
						numShares, (*pricesPtr)[j].Value, value, (*pricesPtr)[j].Time)
				}
			}
			if !foundSellPrice {
				log.Panicln("Unable to find price for sell date: ", (*outcomesPtr).SellDates[i])
			}
		}
	}

	// If we are still in the market, then sell on the last day to value the portfolio
	if (inMarket) {
		value = numShares * (*pricesPtr)[len((*pricesPtr)) - 1].Value
		fmt.Printf("Cashing out: %v Shares sold at %v for %v on %v\n",
			numShares, (*pricesPtr)[len(*pricesPtr) - 1].Value, value, (*pricesPtr)[len(*pricesPtr) - 1].Time)
	}

	(*outcomesPtr).FinalValue = value
	fmt.Printf("Final value: %v\n", (*outcomesPtr).FinalValue)
}
