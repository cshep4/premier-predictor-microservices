package gcp

import (
	"log"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/trace"
)

func Trace() error {
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: "prempred",
	})
	if err != nil {
		log.Printf("error_initialising_tracer: %v", err)
	}
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	return nil
}