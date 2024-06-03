package telemetry

import (
	"context"
	"log"
	"myapp/crossCutting/telemetry/config"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

/*
OpenTelemetry instrumentation is initialized here
Jaeger is set up as a exporter, where all our traces will be exported
Global trace provider is created and set, other libraries or our code can use this in the future to create spans.
*/

func Initialize(cfg config.OpenTelemetryCfg) *tracesdk.TracerProvider {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(cfg.CollectorEndpoint)))
	if err != nil {
		log.Fatalln(err)
	}

	traceProvider := tracesdk.NewTracerProvider(
		// Always be sure to batch in production.
		tracesdk.WithBatcher(exp),
		// Record information about this application in a Resource.
		tracesdk.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(cfg.ServiceName),
			//TODO: dev is correct
			attribute.String("environment", "dev"),
		)),
	)

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it by using `otel.GetTracerProvider()`.
	otel.SetTracerProvider(traceProvider)
	return traceProvider
}

func Terminate(tp *tracesdk.TracerProvider) {
	// Cleanly shutdown and flush telemetry when the application exits.
	exitCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := tp.Shutdown(exitCtx); err != nil {
		log.Println(err.Error())
	}
}
