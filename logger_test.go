package logger

import "testing"

func TestFileLogger(t *testing.T) {
	logger := NewFileLogger(LogLevelDebug, "D:\\logs\\", "test")
	logger.Debug("user id[%d] is come from china ", 32334)
	logger.Warn("test warn log")
	logger.Fatal("test Fatal log")
	logger.Close()
}

func TestConsoleLogger(t *testing.T) {
	logger := NewConsoleLogger(LogLevelDebug)
	logger.Debug("user id[%d] is come from china ", 23333)
	logger.Warn("test console warn log")
	logger.Fatal("test console Fatal log")
	logger.Close()
}
