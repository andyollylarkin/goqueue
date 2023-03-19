package utils

import (
	"github.com/sirupsen/logrus"
	sysloghook "github.com/sirupsen/logrus/hooks/syslog"
	"log/syslog"
)

func NewLogger() *logrus.Logger {
	var l logrus.Logger
	hook, err := sysloghook.NewSyslogHook("", "", syslog.LOG_LOCAL7, "go-queue")
	if err != nil {
		l.Hooks.Add(hook)
	}
	l.Formatter = new(logrus.JSONFormatter)
	return &l
}
