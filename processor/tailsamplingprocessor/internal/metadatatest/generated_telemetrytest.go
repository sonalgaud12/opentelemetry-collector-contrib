// Code generated by mdatagen. DO NOT EDIT.

package metadatatest

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/component/componenttest"
	"go.opentelemetry.io/collector/processor"
	"go.opentelemetry.io/collector/processor/processortest"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/metric/metricdata/metricdatatest"
)

func NewSettings(tt *componenttest.Telemetry) processor.Settings {
	set := processortest.NewNopSettings()
	set.ID = component.NewID(component.MustNewType("tail_sampling"))
	set.TelemetrySettings = tt.NewTelemetrySettings()
	return set
}

func AssertEqualProcessorTailSamplingCountSpansSampled(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_tail_sampling_count_spans_sampled",
		Description: "Count of spans that were sampled or not per sampling policy",
		Unit:        "{spans}",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_processor_tail_sampling_count_spans_sampled")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorTailSamplingCountTracesSampled(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_tail_sampling_count_traces_sampled",
		Description: "Count of traces that were sampled or not per sampling policy",
		Unit:        "{traces}",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_processor_tail_sampling_count_traces_sampled")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorTailSamplingEarlyReleasesFromCacheDecision(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_tail_sampling_early_releases_from_cache_decision",
		Description: "Number of spans that were able to be immediately released due to a decision cache hit.",
		Unit:        "{spans}",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_processor_tail_sampling_early_releases_from_cache_decision")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorTailSamplingGlobalCountTracesSampled(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_tail_sampling_global_count_traces_sampled",
		Description: "Global count of traces that were sampled or not by at least one policy",
		Unit:        "{traces}",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_processor_tail_sampling_global_count_traces_sampled")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorTailSamplingNewTraceIDReceived(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_tail_sampling_new_trace_id_received",
		Description: "Counts the arrival of new traces",
		Unit:        "{traces}",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_processor_tail_sampling_new_trace_id_received")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorTailSamplingSamplingDecisionLatency(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.HistogramDataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_tail_sampling_sampling_decision_latency",
		Description: "Latency (in microseconds) of a given sampling policy",
		Unit:        "µs",
		Data: metricdata.Histogram[int64]{
			Temporality: metricdata.CumulativeTemporality,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_processor_tail_sampling_sampling_decision_latency")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorTailSamplingSamplingDecisionTimerLatency(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.HistogramDataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_tail_sampling_sampling_decision_timer_latency",
		Description: "Latency (in milliseconds) of each run of the sampling decision timer",
		Unit:        "ms",
		Data: metricdata.Histogram[int64]{
			Temporality: metricdata.CumulativeTemporality,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_processor_tail_sampling_sampling_decision_timer_latency")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorTailSamplingSamplingLateSpanAge(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.HistogramDataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_tail_sampling_sampling_late_span_age",
		Description: "Time (in seconds) from the sampling decision was taken and the arrival of a late span",
		Unit:        "s",
		Data: metricdata.Histogram[int64]{
			Temporality: metricdata.CumulativeTemporality,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_processor_tail_sampling_sampling_late_span_age")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorTailSamplingSamplingPolicyEvaluationError(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_tail_sampling_sampling_policy_evaluation_error",
		Description: "Count of sampling policy evaluation errors",
		Unit:        "{errors}",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_processor_tail_sampling_sampling_policy_evaluation_error")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorTailSamplingSamplingTraceDroppedTooEarly(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_tail_sampling_sampling_trace_dropped_too_early",
		Description: "Count of traces that needed to be dropped before the configured wait time",
		Unit:        "{traces}",
		Data: metricdata.Sum[int64]{
			Temporality: metricdata.CumulativeTemporality,
			IsMonotonic: true,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_processor_tail_sampling_sampling_trace_dropped_too_early")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorTailSamplingSamplingTraceRemovalAge(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.HistogramDataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_tail_sampling_sampling_trace_removal_age",
		Description: "Time (in seconds) from arrival of a new trace until its removal from memory",
		Unit:        "s",
		Data: metricdata.Histogram[int64]{
			Temporality: metricdata.CumulativeTemporality,
			DataPoints:  dps,
		},
	}
	got, err := tt.GetMetric("otelcol_processor_tail_sampling_sampling_trace_removal_age")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func AssertEqualProcessorTailSamplingSamplingTracesOnMemory(t *testing.T, tt *componenttest.Telemetry, dps []metricdata.DataPoint[int64], opts ...metricdatatest.Option) {
	want := metricdata.Metrics{
		Name:        "otelcol_processor_tail_sampling_sampling_traces_on_memory",
		Description: "Tracks the number of traces current on memory",
		Unit:        "{traces}",
		Data: metricdata.Gauge[int64]{
			DataPoints: dps,
		},
	}
	got, err := tt.GetMetric("otelcol_processor_tail_sampling_sampling_traces_on_memory")
	require.NoError(t, err)
	metricdatatest.AssertEqual(t, want, got, opts...)
}

func getMetricFromResource(name string, got metricdata.ResourceMetrics) metricdata.Metrics {
	for _, sm := range got.ScopeMetrics {
		for _, m := range sm.Metrics {
			if m.Name == name {
				return m
			}
		}
	}

	return metricdata.Metrics{}
}

func lenMetrics(got metricdata.ResourceMetrics) int {
	metricsCount := 0
	for _, sm := range got.ScopeMetrics {
		metricsCount += len(sm.Metrics)
	}

	return metricsCount
}
