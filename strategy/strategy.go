package strategy

import (
	"time"
	"github.com/cruskit/finsight/metric"
	"github.com/cruskit/finsight/analyser"
	"log"
	"strconv"
)

type StrategyOutcomes struct {
	PurchaseDates [] time.Time `json:"purchaseDates"`
	SellDates     [] time.Time `json:"sellDates"`
	FinalValue    float64 `json:"finalValue"`
	Positions     [] Position `json:"positions"`
	Settings      map[string]string `json:"settings"`
}

type Position struct {
	Date       time.Time `json:"date"`
	NumUnits   float64 `json:"numUnits"`
	Cash       float64 `json:"cash"`
	LastAction string `json:"lastAction"`
}


// START_VALUE is how much money the strategy starts with when tracking trades
var START_VALUE float64 = 1000

func RunBuyAndHold(firstAllowedBuyDate time.Time, pricesPtr *[]metric.Metric) *StrategyOutcomes {

	buy := make([]time.Time, 0)
	sell := make([]time.Time, 0)

	// Work out the date of first purchase
	inMarket := false
	for i := 0 ; i < len(*pricesPtr) && !inMarket ; i++ {
		if firstAllowedBuyDate.Before((*pricesPtr)[i].Time) || firstAllowedBuyDate.Equal((*pricesPtr)[i].Time){
			buy = append(buy, (*pricesPtr)[i].Time)
			inMarket = true;
		}
	}

	outComes := StrategyOutcomes{buy, sell, 0.0, make([]Position, 0), make(map[string]string)}
	calculateStrategyValue(&outComes, pricesPtr)

	// Keep track of what we used
	outComes.Settings["strategy"] = "buyAndHold"
	outComes.Settings["firstAllowedBuy"] = firstAllowedBuyDate.String()

	return &outComes
}

func RunMovingAverageCrossover(fastAvgPeriod int, slowAvgPeriod int, firstAllowedBuyDate time.Time, pricesPtr *[]metric.Metric) *StrategyOutcomes {

	buyDates, sellDates := calculateMovingAverageTxDates(fastAvgPeriod, slowAvgPeriod, firstAllowedBuyDate, pricesPtr)

	outComes := StrategyOutcomes{*buyDates, *sellDates, 0.0, make([]Position, 0), make(map[string]string)}
	calculateStrategyValue(&outComes, pricesPtr)

	// Keep track of what we used
	outComes.Settings["strategy"] = "movingAverageCrossover"
	outComes.Settings["fastAvgPeriod"] = strconv.Itoa(fastAvgPeriod)
	outComes.Settings["slowAvgPeriod"] = strconv.Itoa(slowAvgPeriod)
	outComes.Settings["firstAllowedBuy"] = firstAllowedBuyDate.String()

	return &outComes
}

// calculateMovingAverageTxDates determines the buy and sell dates for a unit based on moving average crossovers
// it returns the buyDates and the sellDates
func calculateMovingAverageTxDates(fastAvgPeriod int, slowAvgPeriod int, firstAllowedBuyDate time.Time, pricesPtr *[]metric.Metric) (*[]time.Time, *[]time.Time) {

	fastAvg := analyser.MovingAverage(fastAvgPeriod, pricesPtr)
	slowAvg := analyser.MovingAverage(slowAvgPeriod, pricesPtr)

	// Make the avgs start on the same date
	fastAvg = fastAvg[slowAvgPeriod - fastAvgPeriod:]

	// Work out the buy and sell dates
	buy := make([]time.Time, 0)
	sell := make([]time.Time, 0)

	for i := 1; i < len(fastAvg); i++ {

		if fastAvg[i].Value > slowAvg[i].Value &&
		fastAvg[i - 1].Value < slowAvg[i - 1].Value &&
		( firstAllowedBuyDate.Before(fastAvg[i].Time) || firstAllowedBuyDate.Equal(fastAvg[i].Time) ) {
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

	var value float64 = START_VALUE
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
//				fmt.Printf("%v units purchased at %v for %v on %v\n",
//					numShares, (*pricesPtr)[j].Value, value, (*pricesPtr)[j].Time)
				(*outcomesPtr).Positions = append((*outcomesPtr).Positions,
					Position{(*pricesPtr)[j].Time, numShares, 0, "buy"})
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
//					fmt.Printf("%v units sold at %v for %v on %v\n",
//						numShares, (*pricesPtr)[j].Value, value, (*pricesPtr)[j].Time)
					(*outcomesPtr).Positions = append((*outcomesPtr).Positions,
						Position{(*pricesPtr)[j].Time, 0, value, "sell"})
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
//		fmt.Printf("Cashing out: %v Shares sold at %v for %v on %v\n",
//			numShares, (*pricesPtr)[len(*pricesPtr) - 1].Value, value, (*pricesPtr)[len(*pricesPtr) - 1].Time)
	}

	(*outcomesPtr).FinalValue = value
//	fmt.Printf("Final value: %v\n", (*outcomesPtr).FinalValue)
}

func CalculateDailyPortfolioValues(outcomes *StrategyOutcomes, pricesPtr *[]metric.Metric) *[] float64 {

	var valuation [] float64
	var currentPosition int = 0

	for _, unitPrice := range(*pricesPtr){

		if (currentPosition < (len(outcomes.Positions) - 1) && unitPrice.Time.Equal(outcomes.Positions[currentPosition+1].Date)){

			currentPosition++
		}

		// Value the portfolio
		if (currentPosition == 0 && unitPrice.Time.Before(outcomes.Positions[currentPosition].Date)){
			valuation = append(valuation, START_VALUE)
		}else{
			var val = outcomes.Positions[currentPosition].Cash + (outcomes.Positions[currentPosition].NumUnits * unitPrice.Value)
			valuation = append(valuation, val)
		}


	}
	return &valuation;
}