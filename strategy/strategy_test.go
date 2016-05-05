package strategy

import (

	"testing"
	"time"
	"github.com/cruskit/finsight/metric"
	"reflect"
)

func TestBuyAndHoldWithZeroStartDate(t *testing.T){

	prices := []metric.Metric{
		metric.Metric{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), 2}, // initial buy
		metric.Metric{time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC), 3},
		metric.Metric{time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC), 4},
		metric.Metric{time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC), 5},
		metric.Metric{time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC), 6},
		metric.Metric{time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC), 7},
		metric.Metric{time.Date(2016, time.January, 7, 0, 0, 0, 0, time.UTC), 8}, // close out
	}
	buys := []time.Time{
		time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
	}
	sells := make([]time.Time, 0)
	positions := []Position{
		Position{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), 500, 0, "buy"},
	}

	so := RunBuyAndHold(time.Time{}, &prices)

	if so.FinalValue != 4000 {
		t.Errorf("Expected final value: %v, got %v", 4000, so.FinalValue)
	}

	if (!reflect.DeepEqual(buys, so.PurchaseDates)){
		t.Errorf("Expected buy dates: %v, got %v", buys, so.PurchaseDates)
	}

	if (!reflect.DeepEqual(sells, so.SellDates)){
		t.Errorf("Expected sell dates: %v, got %v", sells, so.SellDates)
	}
	if (!reflect.DeepEqual(positions, so.Positions)){
		t.Errorf("Expected positions: %v, got %v", positions, so.Positions)
	}

}


func TestBuyAndHoldWithSpecificStartDate(t *testing.T){

	prices := []metric.Metric{
		metric.Metric{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), 2},
		metric.Metric{time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC), 3},
		metric.Metric{time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC), 4},// initial buy
		metric.Metric{time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC), 5},
		metric.Metric{time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC), 6},
		metric.Metric{time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC), 7},
		metric.Metric{time.Date(2016, time.January, 7, 0, 0, 0, 0, time.UTC), 8}, // close out
	}
	buys := []time.Time{
		time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC),
	}
	sells := make([]time.Time, 0)
	positions := []Position{
		Position{time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC), 250, 0, "buy"},
	}

	so := RunBuyAndHold(time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC), &prices)

	if so.FinalValue != 2000 {
		t.Errorf("Expected final value: %v, got %v", 2000, so.FinalValue)
	}

	if (!reflect.DeepEqual(buys, so.PurchaseDates)){
		t.Errorf("Expected buy dates: %v, got %v", buys, so.PurchaseDates)
	}

	if (!reflect.DeepEqual(sells, so.SellDates)){
		t.Errorf("Expected sell dates: %v, got %v", sells, so.SellDates)
	}
	if (!reflect.DeepEqual(positions, so.Positions)){
		t.Errorf("Expected positions: %v, got %v", positions, so.Positions)
	}

}

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
		Position{time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC), 500, 0, "buy"},
		Position{time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC), 0, 2000, "sell"},
		Position{time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC), 200, 0, "buy"},
		Position{time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC), 0, 1600, "sell"},
	}

	so := StrategyOutcomes{buys, sells, 0, make([]Position, 0), make(map[string]string)}
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
		metric.Metric{time.Date(2016, time.January, 7, 0, 0, 0, 0, time.UTC), 8}, // final value , but not sold $1600
	}

	positions := []Position{
		Position{time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC), 500, 0, "buy"},
		Position{time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC), 0, 2000, "sell"},
		Position{time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC), 200, 0, "buy"},
	}

	so := StrategyOutcomes{buys, sells, 0, make([]Position, 0), make(map[string]string)}
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

	buys, sells := calculateMovingAverageTxDates(2, 3, time.Time{}, &prices)
	if (len(*buys) > 0 || len(*sells) > 0){
		t.Error("Expected no transactions, got: ", buys, sells)
	}
}

func TestMovingAverageBuySellTxDates(t *testing.T) {

	prices := []metric.Metric{
		metric.Metric{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), 3},
		metric.Metric{time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC), 2},
		metric.Metric{time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC), 1},
		metric.Metric{time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC), 4}, // buy - before date
		metric.Metric{time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC), 1},
		metric.Metric{time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC), 4}, // sell
		metric.Metric{time.Date(2016, time.January, 7, 0, 0, 0, 0, time.UTC), 1}, // buy
		metric.Metric{time.Date(2016, time.January, 8, 0, 0, 0, 0, time.UTC), 4}, // sell
		metric.Metric{time.Date(2016, time.January, 9, 0, 0, 0, 0, time.UTC), 1}, // buy
	}

	expectedBuys := []time.Time{
		time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC),
		time.Date(2016, time.January, 7, 0, 0, 0, 0, time.UTC),
		time.Date(2016, time.January, 9, 0, 0, 0, 0, time.UTC),
	}
	expectedSells := []time.Time{
		time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC),
		time.Date(2016, time.January, 8, 0, 0, 0, 0, time.UTC),
	}


	buys, sells := calculateMovingAverageTxDates(2, 3, time.Time{}, &prices)
	if (!reflect.DeepEqual(expectedBuys, *buys)){
		t.Errorf("Expected buys of \n%v, \ngot: \n%v", expectedBuys, buys)
	}
	if (!reflect.DeepEqual(expectedSells, *sells)){
		t.Errorf("Expected sells of \n%v, \ngot: \n%v", expectedSells, sells)
	}
}


func TestMovingAverageBuySellWithStartDate(t *testing.T) {

	prices := []metric.Metric{
		metric.Metric{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), 3},
		metric.Metric{time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC), 2},
		metric.Metric{time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC), 1},
		metric.Metric{time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC), 4}, // buy - before date
		metric.Metric{time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC), 1},
		metric.Metric{time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC), 4}, // sell
		metric.Metric{time.Date(2016, time.January, 7, 0, 0, 0, 0, time.UTC), 1}, // buy
		metric.Metric{time.Date(2016, time.January, 8, 0, 0, 0, 0, time.UTC), 4}, // sell
		metric.Metric{time.Date(2016, time.January, 9, 0, 0, 0, 0, time.UTC), 1}, // buy
	}

	expectedBuys := []time.Time{
		time.Date(2016, time.January, 7, 0, 0, 0, 0, time.UTC),
		time.Date(2016, time.January, 9, 0, 0, 0, 0, time.UTC),
	}
	expectedSells := []time.Time{
		time.Date(2016, time.January, 8, 0, 0, 0, 0, time.UTC),
	}


	buys, sells := calculateMovingAverageTxDates(2, 3, time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC), &prices)
	if (!reflect.DeepEqual(expectedBuys, *buys)){
		t.Errorf("Expected buys of \n%v, \ngot: \n%v", expectedBuys, buys)
	}
	if (!reflect.DeepEqual(expectedSells, *sells)){
		t.Errorf("Expected sells of \n%v, \ngot: \n%v", expectedSells, sells)
	}

}

func TestCalculateDailyPortfolioValues(t *testing.T){

	prices := []metric.Metric{
		metric.Metric{time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), 0.01},
		metric.Metric{time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC), 2}, // 500 shares
		metric.Metric{time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC), 0.01},
		metric.Metric{time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC), 4}, // $2000
		metric.Metric{time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC), 11}, // 200 share
		metric.Metric{time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC), 8}, // $1600
		metric.Metric{time.Date(2016, time.January, 7, 0, 0, 0, 0, time.UTC), 0.01},
	}

	positions := []Position{
		Position{time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC), 500, 0, "buy"},
		Position{time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC), 0, 2000, "sell"},
		Position{time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC), 200, 0, "buy"},
		Position{time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC), 0, 1600, "sell"},
	}

	outcomes := StrategyOutcomes{}
	outcomes.Positions = positions

	valuations := CalculateDailyPortfolioValues(&outcomes, &prices)

	expectedValuations := [] float64 {
		START_VALUE,
		1000.00,
		5.0,
		2000.0,
		2200.0,
		1600.0,
		1600.0,
	}

	if (!reflect.DeepEqual(expectedValuations, *valuations)){
		t.Errorf("Expected valuations of \n%v, \ngot: \n%v", expectedValuations, *valuations)
	}

}
