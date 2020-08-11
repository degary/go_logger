package logger

import (
	"fmt"
	"os"
)

//2018/03/26 0:01.386 DEBUG logDebug.go:29 this is a debug log
type FileLogger struct {
	level    int
	logPath  string
	logName  string
	file     *os.File
	warnFile *os.File
}

func NewFileLogger(config map[string]string) (logger LogInterface, err error) {
	logPath, ok := config["log_path"]
	if !ok {
		err = fmt.Errorf("not found log_path")
		return
	}

	logName, ok := config["log_name"]
	if !ok {
		err = fmt.Errorf("not found log_name")
		return
	}

	logLevel, ok := config["log_level"]
	if !ok {
		err = fmt.Errorf("not found log_level")
		return
	}

	level := getLogLevel(logLevel)

	logger = &FileLogger{
		level:   level,
		logPath: logPath,
		logName: logName,
	}
	logger.init()
	return
}

func (f *FileLogger) init() {
	fileName := fmt.Sprintf("%s/%s.log", f.logPath, f.logName)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %v failed,error: %v", file, err))
	}
	f.file = file

	//写错误日志和fatal日志的文件
	fileName = fmt.Sprintf("%s/%s.log.wf", f.logPath, f.logName)
	file, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("open file %s failed,error: %v", fileName, err))
	}

	f.warnFile = file
}

func (f *FileLogger) SetLevel(level int) {
	if level < LogLevelDebug || level > LogLevelFatal {
		level = LogLevelDebug
	}
	f.level = level
}

func (f *FileLogger) Debug(format string, args ...interface{}) {
	if f.level > LogLevelDebug {
		return
	}
	writeLog(f.file, LogLevelDebug, format, args...)
}

func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	writeLog(f.file, LogLevelTrace, format, args...)

}

func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	writeLog(f.file, LogLevelInfo, format, args...)

}

func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	writeLog(f.warnFile, LogLevelWarn, format, args...)
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	writeLog(f.warnFile, LogLevelError, format, args...)
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	writeLog(f.warnFile, LogLevelFatal, format, args...)
}

func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}
