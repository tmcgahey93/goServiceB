package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var requestCount = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "myapp_http_requests_total",
		Help: "Total number of HTTP requests",
	},
)

func init() {
	prometheus.MustRegister(requestCount)
}

func main() {

	http.HandleFunc("/", serviceBHandler)

	// Register /metrics handler to expose Prometheus metrics
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Server Running on port 8081....")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func serviceBHandler(w http.ResponseWriter, r *http.Request) {
	returnMessage := "Hello from Service B"
	fmt.Fprintln(w, returnMessage)
	requestCount.Inc() // increment by 1
}
