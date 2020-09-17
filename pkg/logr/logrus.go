package logr

import (
	"github.com/go-logr/logr"
	"github.com/sirupsen/logrus"
)

type LogrusLogr struct{
	*logrus.Logger
	fields logrus.Fields
}

// Rewrite logr logs to Logrus logger
func Logrus() *LogrusLogr{
	return &LogrusLogr{
		Logger: logrus.StandardLogger(),
		fields: logrus.Fields{},
	}
}

func (log *LogrusLogr) Enabled() bool {
	return true
}

func (log *LogrusLogr) Info(msg string, keysAndValues ...interface{}) {
	fields := logrus.Fields{}

	for k, v := range log.fields {
		fields[k] = v
	}

	for i := 0; i<len(keysAndValues); i+=2 {
		k := keysAndValues[i]
		v := keysAndValues[i+1]
		fields[k.(string)] = v
	}

	log.WithFields(fields).Info(msg)
}

func (log *LogrusLogr) Error(err error, msg string, keysAndValues ...interface{}) {
	fields := logrus.Fields{}

	for k, v := range log.fields {
		fields[k] = v
	}

	for i := 0; i<len(keysAndValues); i+=2 {
		k := keysAndValues[i]
		v := keysAndValues[i+1]
		fields[k.(string)] = v
	}

	log.WithError(err).WithFields(fields).Error(msg)
}

func (log *LogrusLogr) V(level int) logr.Logger {
	fields := logrus.Fields{}

	for k, v := range log.fields {
		fields[k] = v
	}

	fields["level"] = level

	l := &LogrusLogr{
		Logger: log.Logger,
		fields: fields,
	}

	return l
}

// WithValues adds some key-value pairs of context to a logger.
// See Info for documentation on how key/value pairs work.
func (log *LogrusLogr) WithValues(keysAndValues ...interface{}) logr.Logger {
	fields := logrus.Fields{}

	for k, v := range log.fields {
		fields[k] = v
	}

	for i := 0; i<len(keysAndValues); i+=2 {
		k := keysAndValues[i]
		v := keysAndValues[i+1]
		fields[k.(string)] = v
	}

	l:= &LogrusLogr{
		Logger: log.Logger,
		fields: fields,
	}

	return l
}

// WithName adds a new element to the logger's name.
// Successive calls with WithName continue to append
// suffixes to the logger's name.  It's strongly recommended
// that name segments contain only letters, digits, and hyphens
// (see the package documentation for more information).
func (log *LogrusLogr) WithName(name string) logr.Logger {
	l := &LogrusLogr{
		Logger: log.Logger,
		fields: log.fields,
	}

	l.fields["name"] = name

	return l
}
