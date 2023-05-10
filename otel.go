/*
 * Copyright (c) 2023 sixwaaaay.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"context"
	"github.com/sixwaaaay/shauser/internal/config"
	"go.opentelemetry.io/otel/propagation"
	"os"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/runtime"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv/v1.18.0"
)

// TracerProvider constructs a new trace provider.
func TracerProvider(conf *config.Config) (*trace.TracerProvider, error) {
	otel.SetTextMapPropagator(propagation.TraceContext{})
	if conf.Otel.Enabled == false {
		return trace.NewTracerProvider(), nil
	}
	exporter, err := otlptracegrpc.New(context.Background(),
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(conf.Otel.Endpoint),
		otlptracegrpc.WithCompressor("gzip"),
	)
	if err != nil {
		return nil, err
	}
	name, err := os.Hostname()
	if err != nil {
		return nil, err
	}
	tp := trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(
			resource.NewWithAttributes(
				semconv.SchemaURL,
				semconv.ServiceName(conf.Otel.Service),
				semconv.DeploymentEnvironment(conf.Otel.Environment),
				semconv.HostName(name),
			),
		),
	)
	otel.SetTracerProvider(tp)
	return tp, nil
}

// MeterProvider constructs a new meter provider.
func MeterProvider(conf *config.Config) (*metric.MeterProvider, error) {
	if conf.Otel.Enabled == false {
		return metric.NewMeterProvider(), nil
	}
	options := []otlpmetricgrpc.Option{
		otlpmetricgrpc.WithEndpoint(conf.Otel.Endpoint),
		otlpmetricgrpc.WithCompressor("gzip"),
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithTemporalitySelector(preferDeltaTemporalitySelector),
	}
	exporter, err := otlpmetricgrpc.New(context.Background(), options...)
	if err != nil {
		return nil, err
	}
	reader := metric.NewPeriodicReader(
		exporter,
		metric.WithInterval(15*time.Second),
	)
	provider := metric.NewMeterProvider(
		metric.WithReader(reader),
		metric.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(conf.Otel.Service),
			semconv.DeploymentEnvironment(conf.Otel.Environment),
		)))
	global.SetMeterProvider(provider)
	if err := runtime.Start(); err != nil {
		return nil, err
	}
	return provider, nil
}

func preferDeltaTemporalitySelector(kind metric.InstrumentKind) metricdata.Temporality {
	switch kind {
	case metric.InstrumentKindCounter,
		metric.InstrumentKindObservableCounter,
		metric.InstrumentKindHistogram:
		return metricdata.DeltaTemporality
	default:
		return metricdata.CumulativeTemporality
	}
}
