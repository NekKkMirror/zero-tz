package logger

import (
	"github.com/NekKkMirror/zero-tz/internal/pkg/config"
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogger(cfg *config.Config) *logrus.Logger {
	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	return log
}
