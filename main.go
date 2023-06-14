package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var ScrapperFailures int

var (
	listenAddress = flag.String("web.listen-address", ":9888", "Address to listen on for web interface.")
	metricsPath   = flag.String("web.metrics-path", "/metrics", "Path under which to expose metrics.")
)

var (
//	winbgp_static = prometheus.NewGauge(prometheus.GaugeOpts{
//		Name:        "winbgp_static",
//		Help:        "Status of WinBGP service",
//		ConstLabels: prometheus.Labels{"destination": "primary"},
//	})
)

func init() {
	// Metrics have to be registered to be exposed:
	//prometheus.MustRegister(winbgp_static)

	//Create a new instance of the collector and
	//register it with the prometheus client.
	prometheus.MustRegister(newServiceCollector())
	prometheus.MustRegister(newRoutesCollector())
	prometheus.MustRegister(newPeersCollector())
	prometheus.MustRegister(newScrapperCollector())
}

func main() {
	//wmi_query()
	//wmi_query2()
	// Set value (Will be done only once on run)
	//winbgp_static.Set(serviceCheck("w32time"))
	log.Fatal(serverMetrics(*listenAddress, *metricsPath))
}

func serverMetrics(listenAddress, metricsPath string) error {
	http.Handle(metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
            <html>
            <head><title>WinBGP Prometheus Exporter</title></head>
            <body>
            <h1>WinBGP Metrics</h1>
            <p><a href='` + metricsPath + `'>Metrics</a></p>
            </body>
            </html>
        `))
	})
	return http.ListenAndServe(listenAddress, nil)
}
