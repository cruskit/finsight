package main

import (
	"fmt"
	"sort"
	"time"
	"github.com/cruskit/finsight/scenario"
	"github.com/cruskit/finsight/strategy"
	"github.com/cruskit/finsight/restservice"
	"net/http"
	"log"
)

func main() {

	fmt.Println("Hello world")

	//	cpa := []metric.ClosingPrice{
	//		metric.ClosingPrice{time.Now(), 1} ,
	//		metric.ClosingPrice{ time.Now(), 3},
	//	}
	//
	//	sh := metric.StockHistory{cpa}
	//
	//	sh2 := metric.StockHistory{};
	//
	//	fmt.Println("StockHistory:", sh)
	//	fmt.Println("StockHistory2:", sh2)
	//
	//
	//	cp := metric.ClosingPrice{time.Now(), 3}
	//	//cp.Price = 3;
	//	fmt.Println("Closing Price:", cp)

	//sp500Metric := *metric.ReadMetricFromYahooCsv("sampledata/sp500_history.csv")

	//	for _, val := range (sp500Metric){
	//		fmt.Println("Value: ", val)
	//	}

	//strategy.RunMovingAverageCrossover(2, 3, time.Time{}, "sampledata/movingaveragestrategy_test_data.csv")

	runSim := false

	if (runSim) {
		firstAllowedBuy := time.Time{}
		//firstAllowedBuy := time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)

		outComes := scenario.RunMovingAverage(firstAllowedBuy, "sampledata/ivv_history.csv")
		buyAndHoldOutcomes := scenario.RunBuyAndHold(firstAllowedBuy, "sampledata/ivv_history.csv")

		*outComes = append(*outComes, *buyAndHoldOutcomes...)

		sort.Sort(StrategyValueSort(*outComes))

		fmt.Println()
		for _, so := range (*outComes) {
			fmt.Printf("Final value: %v, NumTrades, %v, Settings: %v\n", so.FinalValue, len(so.Positions), so.Settings)
		}
	}else {
		router := restservice.NewRouter()
		log.Fatal(http.ListenAndServe(":8080", router))

	}

}

type StrategyValueSort []strategy.StrategyOutcomes

func (m StrategyValueSort) Len() int {
	return len(m)
}
func (m StrategyValueSort) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}
func (m StrategyValueSort) Less(i, j int) bool {
	return m[i].FinalValue < m[j].FinalValue
}


