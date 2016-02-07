package scenario
import (
	"github.com/cruskit/finsight/strategy"
	"github.com/cruskit/finsight/metric"
	"time"
	"fmt"
)


func RunMovingAverage(datafile string) {

	outComes := make([]strategy.StrategyOutcomes, 0)

	unitPrices := metric.ReadMetricFromYahooCsv(datafile)

	for firstAvg := 50; firstAvg < 51; firstAvg ++ {

		for secondAvg := 100; secondAvg < 400; secondAvg += 50 {

			so := strategy.RunMovingAverageCrossover(firstAvg, secondAvg, time.Time{}, unitPrices)
			outComes = append(outComes, *so)
		}

	}

	fmt.Println()
	for _, so := range (outComes){
		fmt.Printf("Final value: %v, NumTrades, %v, Settings: %v\n", so.FinalValue, len(so.Positions), so.Settings)
	}
}

func RunBuyAndHold(datafile string) {
	outComes := make([]strategy.StrategyOutcomes, 0)
	unitPrices := metric.ReadMetricFromYahooCsv(datafile)

	so := strategy.RunBuyAndHold(time.Time{}, unitPrices)
	outComes = append(outComes, *so)

	fmt.Println()
	for _, so := range (outComes){
		fmt.Printf("Final value: %v, NumTrades, %v, Settings: %v\n", so.FinalValue, len(so.Positions), so.Settings)
	}

}




