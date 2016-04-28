package restservice


import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/cruskit/finsight/strategy"
	"github.com/cruskit/finsight/metric"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func RunStrategy(w http.ResponseWriter, r *http.Request) {

	fastAvg := 241
	slowAvg := 316
	firstAllowedBuyDate := time.Date(2010, time.January, 1, 0, 0, 0, 0, time.UTC)
	unitPrices := metric.ReadMetricFromYahooCsv("sampledata/ivv_history.csv")

	outcome := strategy.RunMovingAverageCrossover(fastAvg, slowAvg, firstAllowedBuyDate, unitPrices)

	if err := json.NewEncoder(w).Encode(outcome); err != nil {
		panic(err)
	}
}

type Todo struct {
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

type Todos []Todo