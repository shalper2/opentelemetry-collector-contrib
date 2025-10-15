// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package windowsservicereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsservicereceiver"

import (
	"golang.org/x/sys/windows"
	"golang.org/x/sys/windows/svc/mgr"
)

/*
* Functions and structs which are used to interact with the Windows service api.
*
* Primary functions are connecting to the Service Control Manager (SCM) to gather service information on scrape.
*
* Docs may be found at https://learn.microsoft.com/en-us/windows/win32/services/services
* and "https://learn.microsoft.com/en-us/windows/win32/api/winsvc/"
**/

// service manager "client"
type serviceManager struct {
	mgr *mgr.Mgr // handle to SCM database
}

// get a connection to the service manager. this modifies the default mgr
// behavior by obtaining a handle with GENERIC_READ permissions instead of the
// default provided by the windows module
func scmConnect() (*serviceManager, error) {
	var m *mgr.Mgr
	var s *uint16

	// we connect with a less permissive generic_read access rather than the
	// default behavior of the mgr.Connect()
	h, err := windows.OpenSCManager(s, nil, windows.GENERIC_READ)
	if err != nil {
		return nil, err
	}

	m.Handle = h

	return &serviceManager{
		m,
	}, nil
}

// disconnect from service manager
func (s *serviceManager) disconnect() error {
	return s.mgr.Disconnect()
}

// get a list of all services the user running the collector process has access
// to
func (s *serviceManager) listServices() ([]string, error) {
	ls, err := s.mgr.ListServices()
	if err != nil {
		return nil, err
	}
	return ls, nil
}
