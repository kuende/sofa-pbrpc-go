package sofa

import "github.com/Sirupsen/logrus"

// Logger defines standard interface for logging
type Logger interface {
	Debug(...interface{})
	Debugln(...interface{})
	Debugf(string, ...interface{})

	Info(...interface{})
	Infoln(...interface{})
	Infof(string, ...interface{})

	Warn(...interface{})
	Warnln(...interface{})
	Warnf(string, ...interface{})

	Error(...interface{})
	Errorln(...interface{})
	Errorf(string, ...interface{})

	Fatal(...interface{})
	Fatalln(...interface{})
	Fatalf(string, ...interface{})

	With(key string, value interface{}) Logger
}

// DefaultLogger implements a default logger
type DefaultLogger struct {
	entry *logrus.Entry
}

// NewLogger returns a new default logger
func NewLogger() Logger {
	return DefaultLogger{entry: logrus.NewEntry(logrus.New())}
}

// With returns a new logger with key=value set
func (l DefaultLogger) With(key string, value interface{}) Logger {
	return DefaultLogger{l.entry.WithField(key, value)}
}

// Debug logs a message at level Debug on the standard logger.
func (l DefaultLogger) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

// Debugln logs a message at level Debug on the standard logger.
func (l DefaultLogger) Debugln(args ...interface{}) {
	l.entry.Debugln(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func (l DefaultLogger) Debugf(format string, args ...interface{}) {
	l.entry.Debugf(format, args...)
}

// Info logs a message at level Info on the standard logger.
func (l DefaultLogger) Info(args ...interface{}) {
	l.entry.Info(args...)
}

// Infoln logs a message at level Info on the standard logger.
func (l DefaultLogger) Infoln(args ...interface{}) {
	l.entry.Infoln(args...)
}

// Infof logs a message at level Info on the standard logger.
func (l DefaultLogger) Infof(format string, args ...interface{}) {
	l.entry.Infof(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func (l DefaultLogger) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func (l DefaultLogger) Warnln(args ...interface{}) {
	l.entry.Warnln(args...)
}

// Warnf logs a message at level Warn on the standard logger.
func (l DefaultLogger) Warnf(format string, args ...interface{}) {
	l.entry.Warnf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func (l DefaultLogger) Error(args ...interface{}) {
	l.entry.Error(args...)
}

// Errorln logs a message at level Error on the standard logger.
func (l DefaultLogger) Errorln(args ...interface{}) {
	l.entry.Errorln(args...)
}

// Errorf logs a message at level Error on the standard logger.
func (l DefaultLogger) Errorf(format string, args ...interface{}) {
	l.entry.Errorf(format, args...)
}

// Fatal logs a message at level Fatal on the standard logger.
func (l DefaultLogger) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

// Fatalln logs a message at level Fatal on the standard logger.
func (l DefaultLogger) Fatalln(args ...interface{}) {
	l.entry.Fatalln(args...)
}

// Fatalf logs a message at level Fatal on the standard logger.
func (l DefaultLogger) Fatalf(format string, args ...interface{}) {
	l.entry.Fatalf(format, args...)
}
