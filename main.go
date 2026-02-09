package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "myapp_http_requests_total",
			Help: "Total number of HTTP requests",
		},
	)

	cpuUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cpu_usage",
			Help: "Current CPU usage",
		},
		[]string{"instance", "pod"},
	)
)

type CpuUsage struct {
	Instance string
	Pod      string
	Usage    float64
}

func init() {
	// Register metric with Prometheus
	prometheus.MustRegister(httpRequestsTotal)
	prometheus.MustRegister(cpuUsage)
}

func updateRequestCount(w http.ResponseWriter, r *http.Request) {
	httpRequestsTotal.Inc() // increment by 1
	w.Write([]byte("Hello, Prometheus!"))
}

func updateCpuUsage(w http.ResponseWriter, r *http.Request) {
	// Simulate CPU usage update
	var cpuUsageBody CpuUsage
	if err := json.NewDecoder(r.Body).Decode(&cpuUsageBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println("Received CPU usage update:", cpuUsageBody)
	cpuUsage.WithLabelValues(cpuUsageBody.Instance, cpuUsageBody.Pod).Set(cpuUsageBody.Usage)
	w.Write([]byte("CPU usage updated!"))
}

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/update-cpu-usage", updateCpuUsage)
	http.HandleFunc("/", updateRequestCount)

	http.ListenAndServe(":8080", nil)
}

// kubestate-metrics
