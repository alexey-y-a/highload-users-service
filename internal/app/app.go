package app

import (
	"context"
	"errors"
	"fmt"
	"net"
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

	log.WithFields(logger.Fields{
		"addr": cfg.GetServerAddr(),
	}).Info("application starting")

	server := &http.Server{
		Addr:              cfg.GetServerAddr(),
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	err = server.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("http server: %w", err)
	}

	return nil
}

func healthzHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ok"}`))
}
