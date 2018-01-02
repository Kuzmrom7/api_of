package log

import (
	"github.com/sirupsen/logrus"
)

// Log describe logger config
type Log struct {
	log     *logrus.Logger
	verbose bool
}

var l *Log

// New initialize logger obj
func New(verbose bool) {
	l = new(Log)
	l.log = logrus.New()
	l.log.SetLevel(logrus.DebugLevel)
	l.verbose = verbose
}

// Debugf print formatted debug log
func Debugf(format string, args ...interface{}) {
	if l.verbose {
		l.log.Debugf(format, args...)
	}
}

// Panicf print formatted panic log
func Panicf(format string, args ...interface{}) {
	l.log.Panicf(format, args...)
}

// Errorf print formatted error log
func Errorf(format string, args ...interface{}) {
	if l.verbose {
		l.log.Errorf(format, args...)
	}
}