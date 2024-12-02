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
	tb.ProcessorGroupbyattrsLogGroups.Record(context.Background(), 1)
	tb.ProcessorGroupbyattrsMetricGroups.Record(context.Background(), 1)
	tb.ProcessorGroupbyattrsNumGroupedLogs.Add(context.Background(), 1)
	tb.ProcessorGroupbyattrsNumGroupedMetrics.Add(context.Background(), 1)
	tb.ProcessorGroupbyattrsNumGroupedSpans.Add(context.Background(), 1)
	tb.ProcessorGroupbyattrsNumNonGroupedLogs.Add(context.Background(), 1)
	tb.ProcessorGroupbyattrsNumNonGroupedMetrics.Add(context.Background(), 1)
	tb.ProcessorGroupbyattrsNumNonGroupedSpans.Add(context.Background(), 1)
	tb.ProcessorGroupbyattrsSpanGroups.Record(context.Background(), 1)
	AssertEqualProcessorGroupbyattrsLogGroups(t, testTel,
		[]metricdata.HistogramDataPoint[int64]{{}}, metricdatatest.IgnoreValue(),
		metricdatatest.IgnoreTimestamp())
	AssertEqualProcessorGroupbyattrsMetricGroups(t, testTel,
		[]metricdata.HistogramDataPoint[int64]{{}}, metricdatatest.IgnoreValue(),
		metricdatatest.IgnoreTimestamp())
	AssertEqualProcessorGroupbyattrsNumGroupedLogs(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualProcessorGroupbyattrsNumGroupedMetrics(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualProcessorGroupbyattrsNumGroupedSpans(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualProcessorGroupbyattrsNumNonGroupedLogs(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualProcessorGroupbyattrsNumNonGroupedMetrics(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualProcessorGroupbyattrsNumNonGroupedSpans(t, testTel,
		[]metricdata.DataPoint[int64]{{Value: 1}},
		metricdatatest.IgnoreTimestamp())
	AssertEqualProcessorGroupbyattrsSpanGroups(t, testTel,
		[]metricdata.HistogramDataPoint[int64]{{}}, metricdatatest.IgnoreValue(),
		metricdatatest.IgnoreTimestamp())

	require.NoError(t, testTel.Shutdown(context.Background()))
}
