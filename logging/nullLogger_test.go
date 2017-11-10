package logging

import "testing"

func TestPrint(t *testing.T) {
	NullLogger{}.Print("xxx")
}

func TestPrintf(t *testing.T) {
	NullLogger{}.Printf("%v", "xxx")
}

func TestDebug(t *testing.T) {
	NullLogger{}.Debug("xxx")
}

func TestDebugf(t *testing.T) {
	NullLogger{}.Debugf("%v", "xxx")
}

func TestInfo(t *testing.T) {
	NullLogger{}.Info("xxx")
}

func TestInfof(t *testing.T) {
	NullLogger{}.Infof("%v", "xxx")
}

func TestWarn(t *testing.T) {
	NullLogger{}.Warn("xxx")
}

func TestWarnf(t *testing.T) {
	NullLogger{}.Warnf("%v", "xxx")
}

func TestError(t *testing.T) {
	NullLogger{}.Error("xxx")
}

func TestErrorf(t *testing.T) {
	NullLogger{}.Errorf("%v", "xxx")
}

func TestFatal(t *testing.T) {
	NullLogger{}.Fatal("xxx")
}

func TestFatalf(t *testing.T) {
	NullLogger{}.Fatalf("%v", "xxx")
}

func TestPanic(t *testing.T) {
	NullLogger{}.Panic("xxx")
}

func TestPanicf(t *testing.T) {
	NullLogger{}.Panicf("%v", "xxx")
}
