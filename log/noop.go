package log

// NoopLogger implements Logger. It do nothing.
type NoopLogger struct {
}

// compile check for interface implementation
var _ Logger = (*NoopLogger)(nil)

// Errorf implements Logger.Errorf.
func (NoopLogger) Errorf(arg0 string, args ...interface{}) {
}

// Tracef implements Logger.Tracef.
func (NoopLogger) Tracef(arg0 string, args ...interface{}) {
}
