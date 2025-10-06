package db

import (
	"os"
	"github.com/joho/godotenv"
	"log"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var InfluxURL = "http://100.87.19.228:8086"
var InfluxOrg = "VIniCosta1"
var InfluxBucket = "Bucket1"
var Client influxdb2.Client

func InfluxDBConn() {
	erro := godotenv.Load()
    if erro != nil {
        log.Fatal("Error loading .env file")
    }

	InfluxToken := os.Getenv("INFLUXDB_TOKEN")

	Client = influxdb2.NewClient(InfluxURL, InfluxToken)
	defer Client.Close()
}