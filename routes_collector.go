package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type routeCollector struct {
	routeDesc *prometheus.Desc
}

func newRouteCollector() *routeCollector {
	return &routeCollector{
		routeDesc: prometheus.NewDesc("winbgp_state_route",
			"WinBGP routes status",
			[]string{"name", "network", "family"},
			nil,
		),
	}
}

// Set description on collector
func (collector *routeCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.routeDesc
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
func (collector *routeCollector) Collect(ch chan<- prometheus.Metric) {
	routes := exec_routes()
	for _, route := range routes {
		ch <- prometheus.MustNewConstMetric(collector.routeDesc, prometheus.GaugeValue, routeStatusConverter(route.Status), route.Name, route.Network, "ipv4")
	}
}
