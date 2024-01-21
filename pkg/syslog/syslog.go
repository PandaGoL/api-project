package syslog

import (
	"github.com/PandaGoL/api-project/pkg/options"
	log "github.com/sirupsen/logrus"
)

// InitLog function initialize logger
func InitLog() (err error) {
	log.Info("Start init logger")
	var lvl log.Level
	if lvl, err = log.ParseLevel(options.Get().LogLevel); err != nil {
		return
	}
	log.SetLevel(lvl)

	return
}
