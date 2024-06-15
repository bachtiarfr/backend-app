package logging

import (
	"github.com/sirupsen/logrus"
	"os"
)

func InitLogger() *logrus.Logger {
	log := logrus.New()
	log.Out = os.Stdout
	log.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}

	return log
}
