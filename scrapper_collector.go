package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type scrapperCollector struct {
	scrapperDesc *prometheus.Desc
}

func newScrapperCollector() *scrapperCollector {
	return &scrapperCollector{
		scrapperDesc: prometheus.NewDesc("winbgp_exporter_parse_failures",
			"Number of errors while parsing output",
			nil,
			nil,
		),
	}
}

// Set description on collector
func (collector *scrapperCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.scrapperDesc
}

// Populate collector with metrics
func (collector *scrapperCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(collector.scrapperDesc, prometheus.CounterValue, float64(ScrapperFailures))
}
