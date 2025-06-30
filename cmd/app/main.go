package main

import (
//	"io"
	"net/http"
	"os"
	
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"05_task/internal/generated/wrappers"
	"05_task/internal/monitors"
)

func initPrometheus() {
	reg := prometheus.NewRegistry()
	prometheus.DefaultRegisterer = reg
	prometheus.DefaultGatherer = reg
	
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		_ = http.ListenAndServe(":2112", nil)
	}()
}

func initLogs() monitors.Monitor {
	baseMonitor := &monitors.BasicMonitor{}
	
	// Оборачиваем в обёртки с нужными параметрами
	logMonitor := wrappers.NewMonitorWithLog(
		baseMonitor,
		os.Stdout,  // output writer
		os.Stderr,  // error writer
	)
	
	metricsMonitor := wrappers.NewMonitorWithPrometheus(
		logMonitor,
		"app_",     // префикс метрик
	)
	
	_ = metricsMonitor.Start()
	_ = metricsMonitor.Log("Initializing monitoring system")
	_ = metricsMonitor.Metric("init", 1.0, map[string]string{"status": "started"})
	
	return metricsMonitor
}

func Task05() {
	initPrometheus()
	monitor := initLogs()
	
	_ = monitor.Log("Starting Task05 execution")
	_ = monitor.Metric("temperature", 23.5, map[string]string{"unit": "celsius"})
	_ = monitor.Stop()
}

func main() {
	Task05()
	select {}
}
