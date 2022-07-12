# winbgp_exporter

This is a prometheus exporter for WinBGP. It currently works with the following WinBGP version:

- 1.3.0 (or higher)


## Usage

The exporter listen on the documented port of `9888`.
If you want to enable the exporter inside WinBGP configuration:

```text
"config":
 {
  "PrometheusExporter" : true
 },
```

## metrics

### `winbgp_status`

```text
# HELP winbgp_status Status of WinBGP service
# TYPE winbgp_status gauge
winbgp_status 1
```

WinBGP service status: `1` for up. `0` for down.

### `winbgp_exporter_parse_failures`

```text
# HELP winbgp_exporter_parse_failures number of errors while parsing output
# TYPE winbgp_exporter_parse_failures counter
winbgp_exporter_parse_failures 0
```

Exporter parsing failures

### `winbgp_state_peer`

```text
# HELP winbgp_state_peer WinBGP Peers status
# TYPE winbgp_state_peer gauge
winbgp_state_peer{localip=127.0.0.1",name="router1",peer_asn="64496",peer_ip="127.0.0.1"} 1
```

BGP peers connectivity status: `1` for up. `0` for down, `2` for maintenance.

### `winbgp_state_route`

```text
# HELP winbgp_state_route WinBGP routes status
# TYPE winbgp_state_route gauge
winbgp_state_route{family="ipv4",name="route1",network="192.168.88.0/29"} 0
```

BGP route status:`0` for down, `1` for up, `2` for maintenance.
