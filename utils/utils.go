package utils

import (
	"flag"
	"management/model"
	"os"

	logrus "github.com/sirupsen/logrus"
)

var Logger = logrus.New()

func init() {
	Logger.Out = os.Stdout
}

func SetLogger() {
	logLevel := flag.String(model.LogLevel, model.LogLevelInfo, model.LogLevelDebug)
	flag.Parse()

	switch *logLevel {
	case model.LogLevelDebug:
		Logger.SetLevel(logrus.DebugLevel)
	case model.LogLevelError:
		Logger.SetLevel(logrus.ErrorLevel)
	case model.LogLevelWarning:
		Logger.SetLevel(logrus.WarnLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}
}
