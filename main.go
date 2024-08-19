package main

import (
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	opsRequested.Inc()
	defer opsRequested.Dec()
	// loop
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "go_metrics",
		Subsystem: "prometheus",
		Name:      "processed_record_total",
		Help:      "process metrics count",
	})

	opsRequested = promauto.NewGauge(prometheus.GaugeOpts{
		Namespace: "go_metrics",
		Subsystem: "prometheus",
		Name:      "processed_record_count",
		Help:      "request record count",
	})
)

func main() {
	recordMetrics()
	// metrics reporter
	http.Handle("/metrics", promhttp.Handler())
	err := http.ListenAndServe(os.Getenv("METRICS_PORT"), nil)
	if err != nil {
		panic(err)
	}
}
