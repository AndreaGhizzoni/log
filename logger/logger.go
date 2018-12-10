package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log *logrus.Logger

func GetInstance() *logrus.Logger {
	if log == nil {
		log = logrus.New()
		log.SetFormatter(&logrus.TextFormatter{})
		log.SetOutput(os.Stdout)
	}
	return log
}
