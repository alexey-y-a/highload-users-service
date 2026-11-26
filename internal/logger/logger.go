package logger

import (
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

const (
	serviceName = "highload-users-service"
	version     = "1.0.0"
)

type Fields = logrus.Fields

func New(level, format string) *logrus.Entry {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)

	if strings.EqualFold(format, "json") {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02T15:04:05Z07:00",
		})
	} else {
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp: true,
		})
	}

	parsedLevel, err := logrus.ParseLevel(level)
	if err != nil {
		log.WithError(err).Warnf("unknown log level %q, using info", level)
		parsedLevel = logrus.InfoLevel
	}
	log.SetLevel(parsedLevel)

	return log.WithFields(logrus.Fields{
		"service": serviceName,
		"version": version,
	})
}
