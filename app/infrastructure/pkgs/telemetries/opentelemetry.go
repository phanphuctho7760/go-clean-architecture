package telemetries

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/configs"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/helpers"
)

var (
	intervalLimit = 5 * time.Second
)

type openTelemetry struct {
	span trace.Span
}

// newOpenTelemetry
// This function to create an instance for using - this is can use after init
func newOpenTelemetry() TelemetryItf {
	return &openTelemetry{}
}

// initOpenTelemetry
// This function only to int config open telemetry
func initOpenTelemetry() func() {
	return func() {}
}

// initOpenTelemetryWithOtel
// This function only to int config open telemetry with otel
// This is unnecessary when using sentry with otel
func initOpenTelemetryWithOtel() func() {
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String(configs.ConfigEnvGlobalInstance.GetServiceName()),
		),
	)
	if err != nil {
		log.Fatalf("Fail to to create resource open telemetry. Error: %s at %s\n", err.Error(), helpers.GetCallerLocationSkip(1))
	}

	otelAgentAddr := configs.ConfigEnvGlobalInstance.GetOtelExporterOTLPEndpoint()

	metricExp, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(otelAgentAddr))

	if err != nil {
		log.Fatalf("Failed to create the collector metric exporter. Error: %s at %s\n", err.Error(), helpers.GetCallerLocationSkip(1))
	}

	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(
				metricExp,
				sdkmetric.WithInterval(intervalLimit),
			),
		),
	)
	otel.SetMeterProvider(meterProvider)

	traceClient := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(otelAgentAddr),
		otlptracegrpc.WithDialOption(grpc.WithBlock()))
	traceExp, err := otlptrace.New(ctx, traceClient)
	if err != nil {
		log.Fatalf("Failed to create the collector trace exporter. Error: %s at %s\n", err.Error(), helpers.GetCallerLocationSkip(1))
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExp)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
	)

	// set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetTracerProvider(tracerProvider)

	return func() {
		cxt, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		if err := traceExp.Shutdown(cxt); err != nil {
			otel.Handle(err)
		}
		// pushes any last exports to the receiver
		if err := meterProvider.Shutdown(cxt); err != nil {
			otel.Handle(err)
		}
	}
}

func (receiver *openTelemetry) StartSpan(ctx context.Context, name string) (newCtx context.Context) {
	newCtx, receiver.span = otel.Tracer(configs.ConfigEnvGlobalInstance.GetServiceName()).Start(ctx, name)
	return
}

func (receiver *openTelemetry) EndSpan() {
	receiver.span.End()
}

func (receiver *openTelemetry) GetTraceID() string {
	return receiver.span.SpanContext().TraceID().String()
}

func (receiver *openTelemetry) GetSpanID() string {
	if receiver.span != nil {
		return receiver.span.SpanContext().SpanID().String()
	}
	return "no-active-span"
}

func (receiver *openTelemetry) RecordError(ctx context.Context, err error) {
	receiver.span.RecordError(err)
}

func (receiver *openTelemetry) SetStatusError(ctx context.Context, errStr string) {
	receiver.span.SetStatus(codes.Error, errStr)
}

func (receiver *openTelemetry) SetStatusSuccess(ctx context.Context, message string) {
	receiver.span.SetStatus(codes.Ok, message)
}

func (receiver *openTelemetry) SetStatusUnset(ctx context.Context, message string) {
	receiver.span.SetStatus(codes.Unset, message)
}

func (receiver *openTelemetry) SetAttributeString(ctx context.Context, key string, value string) {
	receiver.span.SetAttributes(attribute.String(key, value))
}
