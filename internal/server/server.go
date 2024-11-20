package server

import (
	"certgen/internal/application"
	"certgen/internal/certgen"
	"certgen/internal/logging"
	"certgen/internal/reqstate"
	"context"
	"errors"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() error {
	app, err := application.Setup()
	if err != nil {
		return err
	}

	// Routing
	router := http.NewServeMux()
	// POST /api/ca
	// GET /api/ca/{caId}/pkey
	// GET /api/ca/{caId}/cert
	// GET /api/ca/{caId}/cert-chain

	// POST /api/ca/{caId}/cert
	// GET /api/ca/{caId}/certs
	// GET /api/ca/{caId}/certs/
	router.HandleFunc("POST /api/ca", certgen.PostCaHandler(app))
	router.HandleFunc("GET /api/ca/{caId}/pkey", certgen.GetCaPkeyHandler(app))
	rootHandler := reqstate.Middleware(logging.Middleware(router))

	// Server Startup
	server := http.Server{
		Addr:         app.Cfg.Address,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler:      rootHandler,
	}

	// Spawns the server in a different goroutine so we use the main goroutine for
	// the graceful shutdown below
	go func() {
		slog.Info("server is running.", "addr", app.Cfg.Address)
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}
			log.Fatalln(err)
		}
	}()

	// Graceful Shutdown
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)
	shutdownSignal := <-signalChannel

	slog.Info("shutdown signal received. starting gracefull shutdown...", "signal", shutdownSignal)

	gracefulShutdownCtx, cancel := context.WithTimeout(context.Background(), app.Cfg.GracefulShutdown)
	defer cancel()

	if err := server.Shutdown(gracefulShutdownCtx); err != nil {
		return fmt.Errorf("graceful shutdown failed. requests or works may broke: %w", err)
	}

	slog.InfoContext(gracefulShutdownCtx, "graceful shutdown succeeded. server is done.")
	return nil
}
