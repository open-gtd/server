package logging

var l Logger = NullLogger{}

func SetLogger(logger Logger) {
	l = logger
}

func GetLogger() Logger {
	return l
}
