package error_tracking

import (
	"context"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
	sentryotel "github.com/getsentry/sentry-go/otel"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	"google.golang.org/grpc"

	"github.com/phanphuctho7760/go-clean-architecture/app/infrastructure/pkgs/configs"
	"github.com/phanphuctho7760/go-clean-architecture/app/utils/helpers"
)

var (
	intervalLimit = 5 * time.Second
)

type sentryErrorTracking struct{}

// newSentryTrackingWithOpenTelemetry
// This function to init sentry and add data of open telemetry when send to sentry
// NOTE This is not send data to OTEL
func newSentryTrackingWithOpenTelemetry() (errorTracking ErrorTrackingItf, errorTrackingShutDownFunc func()) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              configs.ConfigEnvGlobalInstance.GetSentryDNS(),
		EnableTracing:    configs.ConfigEnvGlobalInstance.GetSentryEnableTracing(),
		TracesSampleRate: configs.ConfigEnvGlobalInstance.GetSentryTracesSampleRate(),
		Debug:            configs.ConfigEnvGlobalInstance.GetSentryDebug(),
	})
	if err != nil {
		log.Fatalf("Fail to to init sentry. Error: %s at %s\n", err.Error(), helpers.GetCallerLocationSkip(1))
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(sentryotel.NewSentrySpanProcessor()),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(sentryotel.NewSentryPropagator())

	return &sentryErrorTracking{}, func() {}
}

// newSentryTrackingWithOpenTelemetryOtel
// This function to init sentry and connect and send data to otel
// NOTE this required connect OTEL to send data to OTEL
func newSentryTrackingWithOpenTelemetryOtel() (errorTracking ErrorTrackingItf, errorTrackingShutDownFunc func()) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn:              configs.ConfigEnvGlobalInstance.GetSentryDNS(),
		EnableTracing:    configs.ConfigEnvGlobalInstance.GetSentryEnableTracing(),
		TracesSampleRate: configs.ConfigEnvGlobalInstance.GetSentryTracesSampleRate(),
		Debug:            configs.ConfigEnvGlobalInstance.GetSentryDebug(),
	})
	if err != nil {
		log.Fatalf("Fail to to init sentry. Error: %s at %s\n", err.Error(), helpers.GetCallerLocationSkip(1))
	}

	ctx := context.Background()

	name := configs.ConfigEnvGlobalInstance.GetServiceName()

	res, err := resource.New(ctx,
		resource.WithFromEnv(),
		resource.WithProcess(),
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			// the service name used to display traces in backends
			semconv.ServiceNameKey.String(name),
		),
	)
	if err != nil {
		log.Fatalf("Fail to to create resource sentry. Error: %s at %s\n", err.Error(), helpers.GetCallerLocationSkip(1))
	}

	metricExp, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(configs.ConfigEnvGlobalInstance.GetOtelExporterOTLPEndpoint()))

	if err != nil {
		log.Fatalf("Fail to to create the collector metric exporter sentry. Error: %s at %s\n", err.Error(), helpers.GetCallerLocationSkip(1))
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
		otlptracegrpc.WithEndpoint(configs.ConfigEnvGlobalInstance.GetOtelExporterOTLPEndpoint()),
		otlptracegrpc.WithDialOption(grpc.WithBlock()))
	traceExp, err := otlptrace.New(ctx, traceClient)
	if err != nil {
		log.Fatalf("Fail to to create the collector trace exporter sentry. Error: %s at %s\n", err.Error(), helpers.GetCallerLocationSkip(1))
	}

	bsp := sdktrace.NewBatchSpanProcessor(traceExp)
	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(bsp),
		sdktrace.WithSpanProcessor(sentryotel.NewSentrySpanProcessor()),
	)

	// set global propagator to tracecontext (the default is no-op).
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	otel.SetTextMapPropagator(sentryotel.NewSentryPropagator())
	otel.SetTracerProvider(tracerProvider)

	return &sentryErrorTracking{}, func() {
		cxt, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		if err := traceExp.Shutdown(cxt); err != nil {
			otel.Handle(err)
		}
		// pushes any last exports to the receiver
		if err := meterProvider.Shutdown(cxt); err != nil {
			otel.Handle(err)
		}
		sentry.Flush(2 * time.Second)
	}
}

func (receiver sentryErrorTracking) CaptureError(ctx context.Context, err error) {
	sentry.CaptureException(err)
}

func (receiver sentryErrorTracking) CaptureMessage(ctx context.Context, message string) {
	sentry.CaptureMessage(message)
}

func (receiver sentryErrorTracking) SetContext(ctx context.Context, key, value string) {
	scope := sentry.CurrentHub().Scope()
	scope.SetTag(key, value)
}

func (receiver sentryErrorTracking) SetExtra(ctx context.Context, key string, value any) {
	scope := sentry.CurrentHub().Scope()
	scope.SetExtra(key, value)
}

func (receiver sentryErrorTracking) SetExtras(ctx context.Context, extra map[string]any) {
	scope := sentry.CurrentHub().Scope()
	scope.SetExtras(extra)
}
