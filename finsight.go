package main

import (
	"fmt"
	"github.com/cruskit/finsight/strategy"
	"time"
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

	strategy.RunMovingAverageCrossover(2, 3, time.Time{}, "sampledata/movingaveragestrategy_test_data.csv")
}
