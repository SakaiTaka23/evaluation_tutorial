package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Prometheus counter
var pingCounter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "ping_request_count",
		Help: "No of request handled by Ping handler",
	},
)

func ping(w http.ResponseWriter, req *http.Request) {
	// increase count of the counter
	pingCounter.Inc()
	fmt.Fprintf(w, "pong")
}

func main() {
	// register ping counter
	prometheus.MustRegister(pingCounter)
	http.HandleFunc("/ping", ping)
	// expose the metrics
	http.Handle("/metrics", promhttp.Handler())
	_ = http.ListenAndServe(":8090", nil)
}
