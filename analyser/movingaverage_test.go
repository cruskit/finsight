package analyser

import (
	"fmt"
	"time"
	"testing"
	"github.com/cruskit/finsight/metric"
	"reflect"
)

func TestMovingAvgWithNotEnoughValues(t *testing.T) {

	vals := []metric.Metric{
		metric.Metric{time.Now(), 1.} ,
		metric.Metric{ time.Now(), 3.},
	}

	result := MovingAverage(3, &vals)

	if result != nil {
		t.Error("Expected nil, got", vals)
	}

}

func TestMovingAverageMultipleValues(t *testing.T){

	now := time.Now()

	vals := []metric.Metric{
		metric.Metric{now.AddDate(0,-3,0), 1.} ,
		metric.Metric{now.AddDate(0,-2,0), 2.} ,
		metric.Metric{now.AddDate(0,-1,0), 3.} ,
	}

	result := MovingAverage(2, &vals)

	expectedResult := []metric.Metric{
		metric.Metric{now.AddDate(0,-2,0), 1.5} ,
		metric.Metric{now.AddDate(0,-1,0), 2.5} ,
	}

	if !reflect.DeepEqual(expectedResult, result){
		t.Error("Expected ", expectedResult, " got ", result)
	}

}

func TestMovingAverageWithOneResult(t *testing.T){

	now := time.Now()

	vals := []metric.Metric{
		metric.Metric{now.AddDate(0,-3,0), 1.} ,
		metric.Metric{now.AddDate(0,-2,0), 2.} ,
		metric.Metric{now.AddDate(0,-1,0), 3.} ,
	}

	result := MovingAverage(3, &vals)

	expectedResult := []metric.Metric{
		metric.Metric{now.AddDate(0,-1,0), 2.0} ,
	}

	if !reflect.DeepEqual(expectedResult, result){
		t.Error("Expected ", expectedResult, " got ", result)
	}

}

func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
}