package logger

import (
	"bs/configs"
	log "github.com/sirupsen/logrus"
)

func NewLog() *log.Logger {
	l := log.New()
	l.Level = configs.DebugLevel 
	return l
}
