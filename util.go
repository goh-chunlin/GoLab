package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"text/template"
	"time"

	"github.com/Microsoft/ApplicationInsights-Go/appinsights"
)

func checkError(err error) {
	if err != nil {
		client := appinsights.NewTelemetryClient(os.Getenv("APPINSIGHTS_INSTRUMENTATIONKEY"))

		trace := appinsights.NewTraceTelemetry(err.Error(), appinsights.Error)
		trace.Timestamp = time.Now()

		client.Track(trace)

		// // false indicates that we should have this handle the panic, and
		// // not re-throw it.
		// defer appinsights.TrackPanic(client, false)

		panic(err)
	}
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		checkError(err)
	}

}

func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}
