package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type fooCollector struct {
	fooMetric *prometheus.Desc
}

func newFooCollector() *fooCollector {
	return &fooCollector{
		fooMetric: prometheus.NewDesc("winbgp_status",
			"Status of WinBGP service",
			nil, nil,
		),
	}
}

func (collector *fooCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.fooMetric
}

func (collector *fooCollector) Collect(ch chan<- prometheus.Metric) {
	ch <- prometheus.MustNewConstMetric(collector.fooMetric, prometheus.GaugeValue, serviceCheck("w32time"))
}
