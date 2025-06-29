package utils

import (
	"flag"
	"management/model"
	"os"

	logrus "github.com/sirupsen/logrus"
)

var Logger logrus.Logger

func init() {
	Logger = *logrus.New()
	Logger.SetFormatter((&logrus.TextFormatter{
		FullTimestamp: true,
	}))
	Logger.SetOutput(os.Stdout)
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

func Log(logLevel, packageLevel, functionName string, message, parameter interface{}) {
	switch logLevel {
	case model.LogLevelError:
		if parameter != nil {
			Logger.Errorf("packageLevel: %s, functionName: %s, message: %v, parameter: %v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Errorf("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)
		}

	case model.LogLevelWarning:
		if parameter != nil {
			Logger.Warnf("packageLevel: %s, functionName: %s, message: %v, parameter: %v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Warnf("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)
		}

	case model.LogLevelInfo:
		if parameter != nil {
			Logger.Infof("packageLevel: %s, functionName: %s, message: %v, parameter: %v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Infof("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)
		}

	case model.LogLevelDebug:
		if parameter != nil {
			Logger.Debugf("packageLevel: %s, functionName: %s, message: %v, parameter: %v\n", packageLevel, functionName, message, parameter)
		} else {
			Logger.Debugf("packageLevel: %s, functionName: %s, message: %v\n", packageLevel, functionName, message)
		}
	}
}
