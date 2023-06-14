package main

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type routesCollector struct {
	routesDesc *prometheus.Desc
}

var (
	routeStates = []string{
		"up",
		"down",
		"maintenance",
	}
)

func newRoutesCollector() *routesCollector {
	return &routesCollector{
		routesDesc: prometheus.NewDesc("winbgp_state_route",
			"WinBGP routes status",
			[]string{"name", "network", "family", "state"},
			nil,
		),
	}
}

// Set description on collector
func (collector *routesCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.routesDesc
}

// Populate collector with metrics
func (collector *routesCollector) Collect(ch chan<- prometheus.Metric) {
	routes := exec_routes()
	// Parse each routes
	for _, route := range routes {
		// Parse each state
		for _, state := range routeStates {
			isCurrentState := 0.0
			if state == strings.ToLower(route.Status) {
				isCurrentState = 1.0
			}
			ch <- prometheus.MustNewConstMetric(
				collector.routesDesc,
				prometheus.GaugeValue,
				isCurrentState,
				route.Name,
				route.Network,
				"ipv4",
				state,
			)
		}
	}
}
