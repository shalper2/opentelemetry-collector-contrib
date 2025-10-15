// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//revive:disable:unused-parameter
//go:build windows

package windowsservicereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsservicereceiver"

import "golang.org/x/sys/windows/svc/mgr" /**
* Windows Service representation and associated functions. These handle
* interacting with the SCM and martialing service information returned by the
* windows api calls.
**/

// receiver representation of a service
type winService struct {
	name          string
	serviceStatus uint32
	startType     uint32
}

// get a service handler
func getService(mgr *serviceManager, sname string) (*winService, error) {
	var ws *winService
	ws.name = sname

	svc, err := mgr.mgr.OpenService(sname)
	if err != nil {
		return nil, err
	}
	defer svc.Close()

	err = ws.getStatus(svc)
	if err != nil {
		return nil, err
	}

	err = ws.getConfig(svc)
	if err != nil {
		return nil, err
	}

	return ws, nil
}

func (ws *winService) getStatus(svc *mgr.Service) error {
	status, err := svc.Query()
	if err != nil {
		return err
	}
	ws.serviceStatus = uint32(status.State)
	return nil
}

func (ws *winService) getConfig(svc *mgr.Service) error {
	cfg, err := svc.Config()
	if err != nil {
		return err
	}
	ws.startType = cfg.StartType
	return nil
}
