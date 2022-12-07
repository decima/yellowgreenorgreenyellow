package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb-client-go/v2/api"
	log "github.com/sirupsen/logrus"
	"yellowgreenorgreenyellow/config"
)

type StatsToReturn struct {
	Split              map[string]int64   `json:"split"`
	PerLast30Days      map[string][]int64 `json:"per_last_30_days"`
	PerLast30DaysXAxis []string           `json:"per_last_30_days_x_axis"`
	PerMonth           map[string][]int64 `json:"per_month"`
	PerMonthXAxis      []string           `json:"per_month_x_axis"`
}

func influxStats() StatsToReturn {
	res := StatsToReturn{
		Split:              map[string]int64{},
		PerLast30DaysXAxis: []string{},
		PerLast30Days:      map[string][]int64{},
		PerMonthXAxis:      []string{},
		PerMonth:           map[string][]int64{},
	}

	influxClient := initClient()
	_, _, org, bucket := config.InfluxServer()
	queryAPI := influxClient.QueryAPI(org)
	query := `from(bucket: "` + bucket + `")
				|> range(start: 2022-12-06T00:00:00Z)
				|> filter(fn: (r) => r["_measurement"] == "greenyellow" or r["_measurement"] == "yellowgreen")
				|> filter(fn: (r) => r["_field"] == "vote")
				|> cumulativeSum()
				|> last()
				  `
	resultsSplit := runQuery(queryAPI, query)
	for resultsSplit.Next() {
		res.Split[resultsSplit.Record().Measurement()] = resultsSplit.Record().Value().(int64)
	}

	query2 := `from(bucket: "` + bucket + `")
				|> range(start: -30d)
				|> filter(fn: (r) => r["_measurement"] == "greenyellow" or r["_measurement"] == "yellowgreen")
				|> filter(fn: (r) => r["_field"] == "vote")
				|> aggregateWindow(every: 1d, fn: sum, createEmpty: true)
				|> cumulativeSum()`
	resultsPerDay := runQuery(queryAPI, query2)
	processed := map[string]bool{}
	for resultsPerDay.Next() {
		rec := resultsPerDay.Record()
		ts := rec.Time().Format("2006-01-02")
		if _, ok := processed[ts]; !ok {
			res.PerLast30DaysXAxis = append(res.PerLast30DaysXAxis, ts)
			processed[ts] = true
		}
		if _, ok := res.PerLast30Days[rec.Measurement()]; !ok {
			res.PerLast30Days[rec.Measurement()] = []int64{}
		}
		res.PerLast30Days[rec.Measurement()] = append(res.PerLast30Days[rec.Measurement()], rec.Value().(int64))

	}
	res.PerLast30DaysXAxis = append(res.PerLast30DaysXAxis, "current")

	query3 := `from(bucket: "` + bucket + `")
				  |> range(start: 2022-12-01T00:00:00Z)
				  |> filter(fn: (r) => r["_measurement"] == "greenyellow" or r["_measurement"] == "yellowgreen")
				  |> filter(fn: (r) => r["_field"] == "vote")
				  |> aggregateWindow(every: 1mo, fn: sum, createEmpty: true)
				  |> yield(name: "sum")`
	resultsPerMonth := runQuery(queryAPI, query3)

	processed = map[string]bool{}
	for resultsPerMonth.Next() {
		rec := resultsPerMonth.Record()
		ts := rec.Time().Format("Jan 2006")
		if _, ok := processed[ts]; !ok {
			res.PerMonthXAxis = append(res.PerMonthXAxis, ts)
			processed[ts] = true
		}
		if _, ok := res.PerMonth[rec.Measurement()]; !ok {
			res.PerMonth[rec.Measurement()] = []int64{}
		}
		res.PerMonth[rec.Measurement()] = append(res.PerMonth[rec.Measurement()], rec.Value().(int64))

	}
	defer influxClient.Close()

	return res
}

func runQuery(queryAPI api.QueryAPI, query string) *api.QueryTableResult {
	results, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		log.Error(err)
	}
	if err := results.Err(); err != nil {
		log.Error(err)
	}
	return results

}

func Stats(c *gin.Context) {

	c.JSON(200, gin.H{"stats": influxStats()})

}
