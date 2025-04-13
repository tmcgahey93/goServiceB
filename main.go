package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

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
}
