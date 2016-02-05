package strategy

import (

	"testing"
	"time"
	"github.com/cruskit/finsight/metric"
)


func TestCalculateStrategyValueWithEqualBuysAndSells(t *testing.T) {

	buys := []time.Time{
		time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC),
	}

	sells := []time.Time{
		time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC),
	}

	prices := []metric.Metric{
		metric.Metric{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), 0.01},
		metric.Metric{time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC), 2}, // 500 shares
		metric.Metric{time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC), 0.01},
		metric.Metric{time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC), 4}, // $2000
		metric.Metric{time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC), 10}, // 200 share
		metric.Metric{time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC), 8}, // $1600
		metric.Metric{time.Date(2016, time.January, 7, 0, 0, 0, 0, time.UTC), 0.01},
	}

	so := StrategyOutcomes{buys, sells, 0}
	calculateStrategyValue(&so, &prices)

	if (so.FinalValue != 1600.0) {
		t.Errorf("Expected final value of %v, was %v", 1600.0, so.FinalValue)
	}
}

func TestCalculateStrategyValueWithoutFinalSell(t *testing.T) {

	buys := []time.Time{
		time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC),
		time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC),
	}

	sells := []time.Time{
		time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC),
	}

	prices := []metric.Metric{
		metric.Metric{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), 0.01},
		metric.Metric{time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC), 2}, // 500 shares
		metric.Metric{time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC), 0.01},
		metric.Metric{time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC), 4}, // $2000
		metric.Metric{time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC), 10}, // 200 share
		metric.Metric{time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC), 0.01},
		metric.Metric{time.Date(2016, time.January, 7, 0, 0, 0, 0, time.UTC), 8}, // $1600
	}

	so := StrategyOutcomes{buys, sells, 0}
	calculateStrategyValue(&so, &prices)

	if (so.FinalValue != 1600.0) {
		t.Errorf("Expected final value of %v, was %v", 1600.0, so.FinalValue)
	}
}

func TestMovingAverageTxDatesWithNoTransactionResults(t *testing.T){
	// TODO: implement this
}

func TestMovingAverageTxDates(t *testing.T){
	// TODO: implement this
}
