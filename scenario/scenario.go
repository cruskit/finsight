package scenario
import (
	"github.com/cruskit/finsight/strategy"
	"github.com/cruskit/finsight/metric"
	"time"
)


func RunMovingAverage(firstAllowedBuyDate time.Time, datafile string) *[]strategy.StrategyOutcomes {

	outComes := make([]strategy.StrategyOutcomes, 0)

	unitPrices := metric.ReadMetricFromYahooCsv(datafile)

	for firstAvg := 1; firstAvg < 400; firstAvg += 5 {

		for secondAvg := firstAvg + 5; secondAvg < 400; secondAvg += 5 {

			so := strategy.RunMovingAverageCrossover(firstAvg, secondAvg, firstAllowedBuyDate, unitPrices)
			outComes = append(outComes, *so)
		}

	}

	//fmt.Println()
	//for _, so := range (outComes) {
	//	fmt.Printf("Final value: %v, NumTrades, %v, Settings: %v\n", so.FinalValue, len(so.Positions), so.Settings)
	//}

	return &outComes
}

func RunBuyAndHold(firstAllowedBuyDate time.Time, datafile string) *[]strategy.StrategyOutcomes {
	outComes := make([]strategy.StrategyOutcomes, 0)
	unitPrices := metric.ReadMetricFromYahooCsv(datafile)

	so := strategy.RunBuyAndHold(firstAllowedBuyDate, unitPrices)
	outComes = append(outComes, *so)

	//fmt.Println()
	//for _, so := range (outComes) {
	//	fmt.Printf("Final value: %v, NumTrades, %v, Settings: %v\n", so.FinalValue, len(so.Positions), so.Settings)
	//}

	return &outComes
}



