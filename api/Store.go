package api

import (
	"github.com/gin-gonic/gin"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
	"yellowgreenorgreenyellow/config"
)

func writeData(choice string) {
	choice = strings.ToLower(choice)

	influxClient := initClient()
	defer influxClient.Close()
	_, _, org, bucket := config.InfluxServer()

	writeAPI := influxClient.WriteAPI(org, bucket)

	tags := map[string]string{}
	fields := map[string]interface{}{
		"vote": 1,
	}
	point := write.NewPoint(choice, tags, fields, time.Now())
	writeAPI.WritePoint(point)
}

var lockedIP = map[string]time.Time{}

type storeBody struct {
	Choice string `json:"choice"`
}

func Store(c *gin.Context) {
	ip := c.ClientIP()
	log.Warn(ip)
	if t, ok := lockedIP[ip]; ok {
		if t.After(time.Now()) {
			c.JSON(429, gin.H{
				"message": "Seems like you are an extreme voter, please wait " + strconv.Itoa(int(lockedIP[ip].Sub(time.Now()).Seconds())) + " seconds before voting again.",
				"until":   lockedIP[ip].Format("2006-01-02T15-04-05"),
			})
			return
		} else {
			delete(lockedIP, ip)
		}
	}
	body := storeBody{}
	c.Bind(&body)

	writeData(body.Choice)
	lockedIP[ip] = time.Now().Add(time.Minute)
	c.JSON(200, gin.H{"stored": true})

}
