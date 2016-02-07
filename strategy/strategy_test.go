package strategy

import (

	"testing"
	"time"
	"github.com/cruskit/finsight/metric"
	"reflect"
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

	positions := []Position{
		Position{time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC), 500, 0},
		Position{time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC), 0, 2000},
		Position{time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC), 200, 0},
		Position{time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC), 0, 1600},
	}

	so := StrategyOutcomes{buys, sells, 0, make([]Position, 0)}
	calculateStrategyValue(&so, &prices)

	if (so.FinalValue != 1600.0) {
		t.Errorf("Expected final value of %v, was %v", 1600.0, so.FinalValue)
	}
	if (!reflect.DeepEqual(positions, so.Positions)) {
		t.Errorf("Expected positions: \n%v, \nWas: \n%v", positions, so.Positions)
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

	positions := []Position{
		Position{time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC), 500, 0},
		Position{time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC), 0, 2000},
		Position{time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC), 200, 0},
	}

	so := StrategyOutcomes{buys, sells, 0, make([]Position, 0)}
	calculateStrategyValue(&so, &prices)

	if (so.FinalValue != 1600.0) {
		t.Errorf("Expected final value of %v, was %v", 1600.0, so.FinalValue)
	}
	if (!reflect.DeepEqual(positions, so.Positions)) {
		t.Errorf("Expected positions: \n%v, \nWas: \n%v", positions, so.Positions)
	}
}

func TestMovingAverageTxDatesWithNoTransactionResults(t *testing.T) {

	prices := []metric.Metric{
		metric.Metric{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), 1},
		metric.Metric{time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC), 2},
		metric.Metric{time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC), 3},
		metric.Metric{time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC), 4},
		metric.Metric{time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC), 4},
		metric.Metric{time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC), 6},
		metric.Metric{time.Date(2016, time.January, 7, 0, 0, 0, 0, time.UTC), 7},
	}

	buys, sells := calculateMovingAverageTxDates(2, 3, &prices)
	if (len(*buys) > 0 || len(*sells) > 0){
		t.Error("Expected no transactions, got: ", buys, sells)
	}
}

func TestMovingAverageTxDates(t *testing.T) {

	prices := []metric.Metric{
		metric.Metric{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), 3},
		metric.Metric{time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC), 2},
		metric.Metric{time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC), 1},
		metric.Metric{time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC), 5},
		metric.Metric{time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC), 1},
		metric.Metric{time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC), 5},
		metric.Metric{time.Date(2016, time.January, 7, 0, 0, 0, 0, time.UTC), 5},
		metric.Metric{time.Date(2016, time.January, 8, 0, 0, 0, 0, time.UTC), 1},
		metric.Metric{time.Date(2016, time.January, 9, 0, 0, 0, 0, time.UTC), 2},
		metric.Metric{time.Date(2016, time.January, 10, 0, 0, 0, 0, time.UTC), 1},
	}

	expectedBuys := []time.Time{
		time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2016, time.January, 7, 0, 0, 0, 0, time.UTC),
		time.Date(2016, time.January, 10, 0, 0, 0, 0, time.UTC),
	}
	expectedSells := []time.Time{
		time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC),
		time.Date(2016, time.January, 8, 0, 0, 0, 0, time.UTC),
	}


	buys, sells := calculateMovingAverageTxDates(2, 3, &prices)
	if (reflect.DeepEqual(expectedBuys, buys)){
		t.Error("Expected buys of %v, got: %v", expectedBuys, buys)
	}
	if (reflect.DeepEqual(expectedSells, sells)){
		t.Error("Expected sells of %v, got: %v", expectedSells, sells)
	}
}
