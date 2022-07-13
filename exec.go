package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

type Route struct {
	Name    string
	Network string
	Status  string
}

func exec_routes() []Route {
	// Create a new context and add a timeout to it
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // The cancel should be deferred so resources are cleaned up

	cmd := exec.CommandContext(ctx, "powershell", "-nologo", "-noprofile", "WinBGP | Select Name, Network, Status | ConvertTo-JSON")
	out, err := cmd.CombinedOutput()

	// We want to check the context error to see if the timeout was executed.
	// The error returned by cmd.Output() will be OS specific based on what
	// happens when a process is killed.
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Command timed out")
	}

	if err != nil {
		fmt.Println(err)
	}

	// Increment scrapper failures counter
	if (err != nil) || (ctx.Err() == context.DeadlineExceeded) {
		ScrapperFailures++
	}

	var routes []Route
	json.Unmarshal([]byte(out), &routes)
	fmt.Printf("Routes : %+v", routes)

	return routes
}

type Peers struct {
	PeerName           string
	LocalIPAddress     string
	PeerIPAddress      string
	PeerASN            string
	ConnectivityStatus string
}

func exec_peers() []Peers {
	// Create a new context and add a timeout to it
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel() // The cancel should be deferred so resources are cleaned up

	cmd := exec.CommandContext(ctx, "powershell", "-nologo", "-noprofile", "Get-BgpPeer | Select-Object PeerName,LocalIPAddress,PeerIPAddress,PeerASN,@{Label='ConnectivityStatus';Expression={$_.ConnectivityStatus.ToString()}} | ConvertTo-Json")
	out, err := cmd.CombinedOutput()

	// We want to check the context error to see if the timeout was executed.
	// The error returned by cmd.Output() will be OS specific based on what
	// happens when a process is killed.
	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Command timed out")
	}

	if err != nil {
		fmt.Println(err)
	}

	// Increment scrapper failures counter
	if (err != nil) || (ctx.Err() == context.DeadlineExceeded) {
		ScrapperFailures++
	}

	var peers []Peers
	json.Unmarshal([]byte(out), &peers)
	fmt.Printf("Peers : %+v", peers)

	return peers
}
