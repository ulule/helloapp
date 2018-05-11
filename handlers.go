package helloapp

import (
	"net/http"
	"os"

	"github.com/go-chi/render"
)

func healthcheck(w http.ResponseWriter, r *http.Request) {
	headers := map[string]string{}

	for k, v := range r.Header {
		if k == "Cookie" {
			continue
		}

		headers[k] = v[0]
	}

	host, _ := os.Hostname()

	render.DefaultResponder(w, r, render.M{
		"remote_addr": r.RemoteAddr,
		"host":        host,
		"headers":     headers,
		"status":      "OK",
		"version":     Version,
		"revision":    Revision,
		"build_time":  BuildTime,
		"compiler":    Compiler,
	})
}
