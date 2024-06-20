package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/99designs/gqlgen/graphql/introspection"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"github.com/kevinmichaelchen/hasura-n-plus-one-remote-rel/internal/handler/graphql"
	"github.com/kevinmichaelchen/hasura-n-plus-one-remote-rel/internal/handler/graphql/generated"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	h := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: &graphql.Resolver{},
			},
		),
	)

	initOTel()

	r := gin.Default()

	r.Use(otelgin.Middleware("nickname-svc-gin-server"))

	r.POST("/query", func(c *gin.Context) {
		log.Info("Received headers", "headers", c.Request.Header)

		h.ServeHTTP(c.Writer, c.Request)
	})

	err := r.Run(":8081")
	if err != nil {
		panic(err)
	}
}

func initOTel() {
	exporter, err := otlptrace.New(
		context.Background(),
		otlptracehttp.NewClient(
			append(
				[]otlptracehttp.Option{
					otlptracehttp.WithInsecure(),
				},
			)...,
		),
	)
	if err != nil {
		panic(err)
	}

	res, err := resource.New(
		context.Background(),
		resource.WithFromEnv(),
	)
	if err != nil {
		panic(err)
	}

	bsp := sdktrace.NewBatchSpanProcessor(exporter)
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSpanProcessor(bsp),
	)

	otel.SetTracerProvider(tp)

	tmp := propagation.NewCompositeTextMapPropagator(
		// Support the W3C Trace Context format.
		propagation.TraceContext{},
		// Support the W3C Baggage format.
		propagation.Baggage{},
	)

	otel.SetTextMapPropagator(tmp)
}
