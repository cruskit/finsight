package metric

import (
	"fmt"
	"testing"
	"time"
	"sort"
	"reflect"
)

func TestCreateStockHistory(t *testing.T) {

	sh := new(Metric)

	fmt.Println(sh)
}


func TestSortMetrics(t *testing.T) {

	m1 := Metric{time.Now().AddDate(0, 0, -1), 1}
	m2 := Metric{time.Now().AddDate(0, 0, -3), 2}
	m3 := Metric{time.Now().AddDate(0, 0, -2), 3}
	m4 := Metric{time.Now().AddDate(0, 0, -4), 4}

	metrics := []Metric{m1, m2, m3, m4 }

	sort.Sort(MetricArray(metrics))

	expectedResult := []Metric{m4, m2, m3, m1}

	if !reflect.DeepEqual(expectedResult, metrics) {
		t.Error("Expected ", expectedResult, " got ", metrics)
	}
}

func TestReadFromYahooCsv(t *testing.T) {

	metric := *ReadMetricFromYahooCsv("../sampledata/yahoo_test_data.csv")

	d1, _ := time.Parse("2006-01-02", "2016-01-27")
	d2, _ := time.Parse("2006-01-02", "2016-01-28")
	d3, _ := time.Parse("2006-01-02", "2016-01-29")
	d4, _ := time.Parse("2006-01-02", "2016-02-01")

	expected := []Metric{
		Metric{d1, 1882.949951},
		Metric{d2, 1893.359985},
		Metric{d3, 1940.23999},
		Metric{d4, 1939.380005},
	}

	if !reflect.DeepEqual(expected, metric) {
		t.Error("Expected ", expected, " got ", metric)
	}
}
