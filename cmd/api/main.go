package main

import (
	"context"

	"github.com/alexey-y-a/highload-users-service/internal/app"
	"github.com/sirupsen/logrus"
)

func main() {
	err := app.Run(context.Background())
	if err != nil {
		logrus.WithError(err).Fatal("failed to start application")
	}
}
