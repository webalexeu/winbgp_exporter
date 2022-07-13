package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type serviceCollector struct {
	serviceDesc *prometheus.Desc
}

func newServiceCollector() *serviceCollector {
	return &serviceCollector{
		serviceDesc: prometheus.NewDesc("winbgp_status",
			"Status of WinBGP service",
			nil, nil,
		),
	}
}

// Set description on collector
func (collector *serviceCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.serviceDesc
}

// Populate collector with metrics
func (collector *serviceCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(collector.serviceDesc, prometheus.GaugeValue, serviceCheck("winbgp"))
}
