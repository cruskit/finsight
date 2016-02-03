package analyser

import (
	"github.com/cruskit/finsight/metric"
)

func MovingAverage(period int, metricsPtr *[]metric.Metric) []metric.Metric {

	metrics := *metricsPtr

	// Ensure we have enough values to actually average
	if period > len(metrics) {
		return nil
	}

	averages := make([]metric.Metric, len(metrics) - period + 1)

	for i := 0; i < len(metrics) - period + 1; i++ {
		inclVals := metrics[ i : i + period]

		var sum float64
		for _, val := range inclVals {
			sum += val.Value
		}

		// Use the last date in the range as the metric date
		averages[i] = metric.Metric{inclVals[len(inclVals) - 1].Time, sum / float64(period)}
	}

	return averages
}
