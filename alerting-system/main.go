package main

import (
	"fmt"
	"time"
)

type Metric struct {
	Name  string
	Value float64
	Time  time.Time
}

type Alert struct {
	MetricName  string
	Threshold   float64
	ActualValue float64
	AlertTime   time.Time
	Message     string
}

func MonitorMetrics(metric chan<- Metric) {
	metrics := []Metric{
		{Name: "5xx", Value: 2, Time: time.Now()},
		{Name: "4xx", Value: 10, Time: time.Now()},
		{Name: "2xx", Value: 100, Time: time.Now()},
	}

	for _, m := range metrics {
		metric <- m
		time.Sleep(1 * time.Second)
	}
	close(metric)
}

func CheckAlert(metrics <-chan Metric, monitorMetrics chan<- Alert) {
	thresholds := map[string]float64{
		"5xx": 2.0,
		"4xx": 5,
		"2xx": 30,
	}

	for m := range metrics {
		if metricVal, ok := thresholds[m.Name]; ok && m.Value > metricVal {
			alert := Alert{
				MetricName:  m.Name,
				ActualValue: m.Value,
				Threshold:   metricVal,
				AlertTime:   m.Time,
				Message:     fmt.Sprintf("%s exceeded threshold", m.Name),
			}

			monitorMetrics <- alert
			time.Sleep(1 * time.Second)
		}
	}
	close(monitorMetrics)
}

func HandleAlerts(alerts <-chan Alert) {
	for alert := range alerts {
		fmt.Printf("%v\n", alert.Message)
	}
}

func main() {
	metricsChan := make(chan Metric)
	alertChan := make(chan Alert)

	go MonitorMetrics(metricsChan)
	go CheckAlert(metricsChan, alertChan)

	HandleAlerts(alertChan)
}
