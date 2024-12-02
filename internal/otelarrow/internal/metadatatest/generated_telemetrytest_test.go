// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"

	"go.opentelemetry.io/collector/component/componenttest"
)

func TestSetupTelemetry(t *testing.T) {
	testTel := componenttest.NewTelemetry()
	tb, err := metadata.NewTelemetryBuilder(testTel.NewTelemetrySettings())
	require.NoError(t, err)
	defer tb.Shutdown()
	require.NoError(t, tb.RegisterOtelarrowAdmissionInFlightBytesCallback(func(_ context.Context, observer metric.Int64Observer) error {
		observer.Observe(1)
		return nil
	}))
	require.NoError(t, tb.RegisterOtelarrowAdmissionWaitingBytesCallback(func(_ context.Context, observer metric.Int64Observer) error {
		observer.Observe(1)
		return nil
	}))
	AssertEqualOtelarrowAdmissionInFlightBytes(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualOtelarrowAdmissionWaitingBytes(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())

	require.NoError(t, testTel.Shutdown(context.Background()))
}
