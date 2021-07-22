package log

// Logger is an interface of some logger.
// TODO usually i would use struct logger. But it is much more complex abstraction which is out of scope for now.
type Logger interface {
	Errorf(arg0 string, args ...interface{})
	Tracef(arg0 string, args ...interface{})
}
