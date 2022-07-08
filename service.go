package main

import (
	"fmt"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

func serviceCheck(service_name string) float64 {
	m, err := mgr.Connect()
	if err != nil {
		return 0
		fmt.Errorf("failed to connect to service manager: %v", err)
	}
	s, err := m.OpenService(service_name)
	if err != nil {
		return 0
		fmt.Errorf("could not access service: %v", err)
	}
	defer s.Close()
	statusCode, err := s.Query()
	if err != nil {
		return 0
		fmt.Errorf("failed to query to service manager: %v", err)
	}
	switch statusCode.State {
	case svc.Stopped:
		fmt.Printf("%s stopped", err)
		return 0
	case svc.Running:
		fmt.Printf("%s running", err)
		return 1
	}
	return 0
}
