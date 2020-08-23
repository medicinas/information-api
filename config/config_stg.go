// +build stg

package config

const (
	DB_HOST                = "mongodb://127.0.0.1:27017"
	DATABASE_NAME          = "bookmark"
	CONNECTION_POOL        = 5
	API_PORT               = 8080
	PROMETHEUS_PUSHGATEWAY = "http://localhost:9091/"
)
