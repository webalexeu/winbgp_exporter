package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type peersCollector struct {
	peersDesc *prometheus.Desc
}

func newPeersCollector() *peersCollector {
	return &peersCollector{
		peersDesc: prometheus.NewDesc("winbgp_state_peer",
			"WinBGP Peers status",
			[]string{"name", "local_ip", "peer_ip", "peer_asn"},
			nil,
		),
	}
}

// Set description on collector
func (collector *peersCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.peersDesc
}

func peerStatusConverter(connectivity_status string) float64 {
	switch {
	case connectivity_status == "Stopped":
		return 2
	case connectivity_status == "Connected":
		return 1
	}
	return 0
}

// Populate collector with metrics
func (collector *peersCollector) Collect(ch chan<- prometheus.Metric) {
	peers := exec_peers()
	for _, peer := range peers {
		ch <- prometheus.MustNewConstMetric(collector.peersDesc, prometheus.GaugeValue, peerStatusConverter(peer.ConnectivityStatus), peer.PeerName, peer.LocalIPAddress, peer.PeerIPAddress, peer.PeerASN)
	}
}
