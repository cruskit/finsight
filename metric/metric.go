package metric

import (
	"time"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sort"
)


type Metric struct {
	Time  time.Time
	Value float64
}

// ReadMetricFromYahooCsv will parse a csv file that has been exported from
// http://finance.yahoo.com/ and convert it to a Metric. Currently
// just imports the adjustedClosingPrice in the metric.
func ReadMetricFromYahooCsv(fileName string) *[]Metric {

	metric := []Metric{}

	file, err := os.Open(fileName)
	if err != nil {
		// err is printable
		// elements passed are separated by space automatically
		fmt.Println("Error:", err)
		return &metric
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return &metric
	}

	fmt.Println("Num records read: ", len(records))
	metric = make([]Metric, len(records)-1)
	for i, record := range records[1:] {

		// TODO: Read the date from the file
		metric[i].Time, _ = time.Parse("2006-01-02", record[0])

		// Just using closing price at the moment
		// TODO: Make the field to read a method parameter
		metric[i].Value, err = strconv.ParseFloat(record[6], 64)
		if err != nil {
			panic(err)
		}
	}

	// Yahoo files are sorted by date desc
	sort.Sort(MetricArray(metric))

	return &metric
}


type MetricArray []Metric
func (m MetricArray) Len() int { return len(m) }
func (m MetricArray) Swap(i, j int) { m[i], m[j] = m[j], m[i] }
func (m MetricArray) Less(i, j int) bool { return m[i].Time.Before(m[j].Time) }