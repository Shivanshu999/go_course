package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/students-api/internal/config"
)

func main() {

	//load config
	cfg := config.MustLoad()

	fmt.Println("Welcome to students api")

	//setup router
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to students api"))
	})

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}
	slog.Info("starting server", slog.String("address", cfg.Address))

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()

		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown gracefully")

}
