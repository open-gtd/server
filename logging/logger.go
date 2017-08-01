package logging

type Logger interface {
	Print(i ...interface{})
	Printf(format string, args ...interface{})
	Debug(i ...interface{})
	Debugf(format string, args ...interface{})
	Info(i ...interface{})
	Infof(format string, args ...interface{})
	Warn(i ...interface{})
	Warnf(format string, args ...interface{})
	Error(i ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(i ...interface{})
	Fatalf(format string, args ...interface{})
	Panic(i ...interface{})
	Panicf(format string, args ...interface{})
}

type NullLogger struct {
}

func (l NullLogger) Print(i ...interface{})                    {}
func (l NullLogger) Printf(format string, args ...interface{}) {}
func (l NullLogger) Debug(i ...interface{})                    {}
func (l NullLogger) Debugf(format string, args ...interface{}) {}
func (l NullLogger) Info(i ...interface{})                     {}
func (l NullLogger) Infof(format string, args ...interface{})  {}
func (l NullLogger) Warn(i ...interface{})                     {}
func (l NullLogger) Warnf(format string, args ...interface{})  {}
func (l NullLogger) Error(i ...interface{})                    {}
func (l NullLogger) Errorf(format string, args ...interface{}) {}
func (l NullLogger) Fatal(i ...interface{})                    {}
func (l NullLogger) Fatalf(format string, args ...interface{}) {}
func (l NullLogger) Panic(i ...interface{})                    {}
func (l NullLogger) Panicf(format string, args ...interface{}) {}
