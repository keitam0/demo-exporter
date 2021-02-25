package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	ctx := context.Background()

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	sv := &http.Server{
		Handler: mux,
		Addr:    ":" + os.Getenv("PORT"),
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func(ctx context.Context) {
		<-sigs
		log.Println("Shutting down ...")
		sv.Shutdown(ctx)
	}(ctx)

	log.Printf("Listening on %s ...", sv.Addr)
	if err := sv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}
}
