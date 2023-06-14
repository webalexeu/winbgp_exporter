package main

import (
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type peersCollector struct {
	peersDesc *prometheus.Desc
}

var (
	peerStates = []string{
		"stopped",
		"connecting",
		"connected",
	}
)

func newPeersCollector() *peersCollector {
	return &peersCollector{
		peersDesc: prometheus.NewDesc("winbgp_state_peer",
			"WinBGP Peers status",
			[]string{"name", "local_ip", "local_asn", "peer_ip", "peer_asn", "state"},
			nil,
		),
	}
}

// Set description on collector
func (collector *peersCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.peersDesc
}

// Populate collector with metrics
func (collector *peersCollector) Collect(ch chan<- prometheus.Metric) {
	peers := exec_peers()
	// Parse each peers
	for _, peer := range peers {
		// Parse each state
		for _, state := range peerStates {
			isCurrentState := 0.0
			if state == strings.ToLower(peer.ConnectivityStatus) {
				isCurrentState = 1.0
			}
			ch <- prometheus.MustNewConstMetric(
				collector.peersDesc,
				prometheus.GaugeValue,
				isCurrentState,
				peer.PeerName,
				peer.LocalIPAddress,
				peer.LocalASN,
				peer.PeerIPAddress,
				peer.PeerASN,
				state,
			)
		}
	}
}
