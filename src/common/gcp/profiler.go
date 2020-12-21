package gcp

import (
	"log"

	"cloud.google.com/go/profiler"
)

func Profile(service, version string) func() error {
	return func() error {
		err := profiler.Start(profiler.Config{
			Service:        service,
			ServiceVersion: version,
			ProjectID:      "prempred",
		})
		if err != nil {
			log.Printf("error_initialising_profiler: %v", err)
		}

		return nil
	}
}
