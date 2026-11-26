package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/alexey-y-a/highload-users-service/internal/config"
	"github.com/alexey-y-a/highload-users-service/internal/logger"
	"github.com/gorilla/mux"
)

func Run(ctx context.Context) error {
	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	log := logger.New(cfg.GetLogLevel(), cfg.GetLogFormat())

	router := mux.NewRouter()
	router.HandleFunc("/healthz", healthzHandler).Methods(http.MethodGet)

	server := &http.Server{
		Addr:              cfg.GetServerAddr(),
		Handler:           router,
		ReadTimeout:       5 * time.Second,
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	log.WithFields(logger.Fields{
		"addr": cfg.GetServerAddr(),
	}).Info("application starting")

	serverErr := make(chan error, 1)

	go func() {
		serverErr <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErr:
		return normalizeServerError(err)
	case <-ctx.Done():
	}

	log.Info("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = server.Shutdown(shutdownCtx)
	if err != nil {
		return fmt.Errorf("shutdown http server: %w", err)
	}

	err = <-serverErr
	if err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("http server: %w", err)
		}
	}

	log.Info("application stopped")

	return nil
}

func normalizeServerError(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return fmt.Errorf("http server: %w", err)
}

func healthzHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
