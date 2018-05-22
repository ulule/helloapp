package helloapp

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TV4/graceful"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type server struct {
	http.Server
}

func (s *server) Shutdown(ctx context.Context) error {
	fmt.Println("Finished shutdown")

	return nil
}

func Run(port int) error {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Get("/healthz", healthcheck)
	r.Get("/", healthcheck)
	r.Get("/sys/health", healthcheck)

	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	srv := http.Server{
		Addr:    addr,
		Handler: chi.ServerBaseContext(context.Background(), r),
	}

	graceful.LogListenAndServe(&server{srv})

	return nil
}
