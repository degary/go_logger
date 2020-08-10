package logger

import "os"

type ConsoleLogger struct {
	level int
}

func (c *ConsoleLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	c.level = level
}

func (c *ConsoleLogger) Debug(format string, args ...interface{}) {
	writeLog(os.Stdout, LogLevelDebug, format, args...)
}

func (c *ConsoleLogger) Trace(format string, args ...interface{}) {
	writeLog(os.Stdout, LogLevelTrace, format, args...)
}
