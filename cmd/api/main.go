package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/alexey-y-a/highload-users-service/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	err := app.Run(ctx)
	if err != nil {
		logrus.WithError(err).Fatal("application failed")
	}
}
