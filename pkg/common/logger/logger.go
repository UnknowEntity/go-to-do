package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var formatter = &logrus.TextFormatter{
	DisableColors: false,
}

var Log *logrus.Logger = &logrus.Logger{
	Out:       os.Stdout,
	Formatter: formatter,
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.DebugLevel,
}
