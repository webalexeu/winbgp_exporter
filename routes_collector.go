package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type routesCollector struct {
	routesDesc *prometheus.Desc
}

func newRoutesCollector() *routesCollector {
	return &routesCollector{
		routesDesc: prometheus.NewDesc("winbgp_state_route",
			"WinBGP routes status",
			[]string{"name", "network", "family"},
			nil,
		),
	}
}

// Set description on collector
func (collector *routesCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.routesDesc
}

func routeStatusConverter(route_status string) float64 {
	switch {
	case route_status == "down":
		return 0
	case route_status == "up":
		return 1
	}
	return 0
}

// Populate collector with metrics
func (collector *routesCollector) Collect(ch chan<- prometheus.Metric) {
	routes := exec_routes()
	for _, route := range routes {
		ch <- prometheus.MustNewConstMetric(collector.routesDesc, prometheus.GaugeValue, routeStatusConverter(route.Status), route.Name, route.Network, "ipv4")
	}
}
