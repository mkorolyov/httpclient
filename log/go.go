package log

import (
	"log"
	"os"
)

// GoLogger implements Logger. It uses stdlib golang log under the hood.
type GoLogger struct {
	l *log.Logger
}

// NewGoLogger constructs GoLogger which prints to stdout.
func NewGoLogger() *GoLogger {
	return &GoLogger{l: log.New(os.Stdout, "", 0)}
}

// compile check for interface implementation
var _ Logger = (*GoLogger)(nil)

// Errorf implements Logger.Errorf. It adds `ERROR: ` prefix to all messages
func (l *GoLogger) Errorf(arg0 string, args ...interface{}) {
	l.l.Printf("ERROR: "+arg0, args...)
}

// Tracef implements Logger.Tracef. It adds `TRACE: ` prefix to all messages
func (l *GoLogger) Tracef(arg0 string, args ...interface{}) {
	l.l.Printf("TRACE: "+arg0, args...)
}
