package helloapp

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/valve"
)

func Run(port int) error {
	r := chi.NewRouter()
	r.Use(middleware.RealIP)
	r.Get("/healthz", healthcheck)
	r.Get("/", healthcheck)
	r.Get("/sys/health", healthcheck)

	valv := valve.New()
	baseCtx := valv.Context()

	addr := fmt.Sprintf(":%s", strconv.Itoa(port))

	srv := http.Server{
		Addr:    addr,
		Handler: chi.ServerBaseContext(baseCtx, r),
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			fmt.Println("shutting down..")

			// first valv
			valv.Shutdown(20 * time.Second)

			// create context with timeout
			ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
			defer cancel()

			// start http shutdown
			srv.Shutdown(ctx)

			// verify, in worst case call cancel via defer
			select {
			case <-time.After(21 * time.Second):
				fmt.Println("not all connections done")
			case <-ctx.Done():

			}
		}
	}()

	return srv.ListenAndServe()
}
