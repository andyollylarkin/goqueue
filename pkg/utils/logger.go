package utils

import (
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogger(isDebugMode bool) *logrus.Logger {
	var l logrus.Logger
	l.Out = os.Stdout
	if isDebugMode {
		l.Level = logrus.DebugLevel
	} else {
		l.Level = logrus.InfoLevel
	}
	//hook, err := sysloghook.NewSyslogHook("", "", syslog.LOG_LOCAL7, "go-queue")
	//if err != nil {
	//	l.Hooks.Add(hook)
	//}
	l.Formatter = new(logrus.TextFormatter)
	return &l
}
