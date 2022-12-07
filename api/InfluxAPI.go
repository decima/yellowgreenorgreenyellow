package api

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"yellowgreenorgreenyellow/config"
)

func initClient() influxdb2.Client {
	url, token, _,_ := config.InfluxServer()
	return influxdb2.NewClient(url, token)
}
