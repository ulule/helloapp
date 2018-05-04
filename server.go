package helloapp

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

func Run(port int) error {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.Status(r, http.StatusOK)
	})
	r.Get("/sys/health", func(w http.ResponseWriter, r *http.Request) {
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
	})

	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	srv := http.Server{
		Addr:    addr,
		Handler: chi.ServerBaseContext(context.Background(), r),
	}

	return srv.ListenAndServe()
}
