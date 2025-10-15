// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build windows

package windowsservicereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsservicereceiver"

import (
	"context"
	"strings"
	"time"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver"
	"go.opentelemetry.io/collector/scraper/scrapererror"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsservicereceiver/internal/metadata"
)

// the scraper is in charge of pulling service information from the os,
// serializing it, and emitting it as metrics
type windowsServiceScraper struct {
	sm       *serviceManager
	services []string
	settings receiver.Settings
	conf     *Config
	mb       *metadata.MetricsBuilder
}

// creates a new windowsServiceScraper
func newWindowsServiceScraper(settings receiver.Settings, cfg *Config) windowsServiceScraper {
	return windowsServiceScraper{
		settings: settings,
		conf:     cfg,
		mb:       metadata.NewMetricsBuilder(metadata.DefaultMetricsBuilderConfig(), settings),
	}
}

// start the scraper by connecting to the windows service manager and compiling
// a list of services to be monitored
func (w *windowsServiceScraper) start(context.Context, component.Host) (err error) {
	scm, err := scmConnect()
	if err != nil {
		return err
	}
	w.sm = scm

	// services to be monitored is the set made by ALL_SERVICES - EXCLUDE_LIST
	var serviceList []string
	services, err := scm.listServices()
	if err != nil {
		return err
	}

	// trivial case (no exclude we monitor everything)
	if len(w.conf.ExcludeServices) == 0 {
		w.services = services
		return nil
	}

	exMap := make(map[string]bool)
	for _, v := range w.conf.ExcludeServices {
		exMap[strings.ToUpper(v)] = true
	}
	for _, v := range services {
		v = strings.ToUpper(v)
		if _, ok := exMap[v]; !ok { // if the value is not in the exMap we will monitor it
			serviceList = append(serviceList, v)
		}
	}

	w.services = serviceList
	return nil
}

// shutdown will close the connection to the windows service manager
func (w *windowsServiceScraper) shutdown(context.Context) (err error) {
	return w.sm.disconnect()
}

// scrape function gathers and emits the metric data for each service in
// w.services
func (w *windowsServiceScraper) scrape(context.Context) (pmetric.Metrics, error) {
	var si []*winService
	var errs scrapererror.ScrapeErrors

	now := pcommon.NewTimestampFromTime(time.Now())

	for _, v := range w.services {
		ws, err := getService(w.sm, v)
		if err != nil {
			errs.AddPartial(1, err)
			continue
		}
		si = append(si, ws)
	}

	for _, v := range si {
		w.mb.RecordWindowsServiceStatusDataPoint(now, int64(v.serviceStatus), v.name, metadata.AttributeStartupMode(v.startType))
	}

	return w.mb.Emit(), errs.Combine()
}
