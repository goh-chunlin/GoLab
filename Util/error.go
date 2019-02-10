package util

import (
	"os"
	"time"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
)

// CheckError is to trace the error and create panic
func CheckError(err error) {
	if err != nil {
		client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

		trace := appinsights.NewTraceTelemetry(err.Error(), appinsights.Error)
		trace.Timestamp = time.Now()

		client.Track(trace)

		// false indicates that we should have this handle the panic, and
		// not re-throw it.
		defer appinsights.TrackPanic(client, false)

		panic(err)
	}
}
