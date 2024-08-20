package main

import (
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// metrics reporter
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(os.Getenv("METRICS_PORT"), nil)
	if err != nil {
		panic(err)
	}
}
