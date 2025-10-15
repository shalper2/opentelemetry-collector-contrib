// Copyright The OpenTelemetry Authors
// SPDX-License-Identifier: Apache-2.0

//go:build !windows

package windowsservicereceiver

import (
	"context"
	"errors"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/consumer"
	"go.opentelemetry.io/collector/receiver"
)

var errUnsupportedOperatingSystem = errors.New("the Windows service receiver is only supported on Windows operating systems.")

func createMetricsReceiver(
	_ context.Context,
	_ receiver.Settings,
	_ component.Config,
	_ consumer.Metrics,
) (receiver.Metrics, error) {
	return nil, errUnsupportedOperatingSystem
}
