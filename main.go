package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/gorilla/mux"
	memoryPrinter "github.com/metrics-tool/memoryusage"
	metrics "github.com/rcrowley/go-metrics"
)

var registry metrics.Registry
var counter metrics.Counter
var gauge metrics.Gauge
var gaugeFloat64 metrics.GaugeFloat64
var memoryStats runtime.MemStats

var gaugeName string = "memoryAlloc"

func initializeRouter() {
	r := mux.NewRouter()

	fmt.Println(gaugeName , " last value: ", gauge.Value(), " bytes")
	memoryStats = memoryPrinter.PrintMemUsage()
	gauge.Update(int64(memoryStats.Alloc))

	fmt.Println("metrics test server running at :: 9000")
	r.HandleFunc("/api/status", logStatus).Methods("GET")

	fmt.Println(gaugeName , " last value: ", gauge.Value(), " bytes")	
	memoryStats = memoryPrinter.PrintMemUsage()
	gauge.Update(int64(memoryStats.Alloc))
	
	log.Fatal(http.ListenAndServe(":9000", r))
}

func logStatus(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HTTP Status: ", "200")

	fmt.Println(gaugeName , " last value: ", gauge.Value(), " bytes")
	memoryStats = memoryPrinter.PrintMemUsage()
	gauge.Update(int64(memoryStats.Alloc))
}


func metricsInit() {
	registry = metrics.NewRegistry()
	gauge = metrics.NewGauge()

	registry.Register(gaugeName, gauge)
}

func main() {
	metricsInit()
	initializeRouter()
}
