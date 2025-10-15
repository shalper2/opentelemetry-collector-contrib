// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//revive:disable:unused-parameter

package windowsservicereceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsservicereceiver"

import (
	"go.opentelemetry.io/collector/scraper/scraperhelper"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/windowsservicereceiver/internal/metadata"
)

// Config defines configuration for windowsservice receiver.
type Config struct {
	scraperhelper.ControllerConfig `mapstructure:",squash"`
	metadata.MetricsBuilderConfig  `mapstructure:",squash"`
	ExcludeServices                []string `mapstructure:"exclude_services"` // user provided list of services to be excluded (if blank all services on the machine are monitored)
}

// Validate checks the receiver configuration is valid
func (*Config) Validate() error {
	return nil
}
