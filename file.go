package logger

import (
	"fmt"
	"os"
	"strconv"
)

//2018/03/26 0:01.386 DEBUG logDebug.go:29 this is a debug log
type FileLogger struct {
	level    int
	logPath  string
	logName  string
	file     *os.File
	warnFile *os.File
	//chan中存放指针,会提高性能
	LogDataChan chan *LogData
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

	//config中传入chan的size大小,类型为string
	logChanSize, ok := config["log_chan_size"]
	if !ok {
		logChanSize = "50000"
	}
	chanSize, err := strconv.Atoi(logChanSize)
	if err != nil {
		chanSize = 50000
	}

	logger = &FileLogger{
		level:       level,
		logPath:     logPath,
		logName:     logName,
		LogDataChan: make(chan *LogData, chanSize),
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
	go f.writeLogBackground()
}

func (f *FileLogger) writeLogBackground() {
	for logData := range f.LogDataChan {
		var file *os.File = f.file
		if logData.WarnAndFatal {
			file = f.warnFile
		}
		fmt.Fprintf(file, "[%s] %s [%s:%s:%d] %s\n", logData.TimeStr, logData.LevelStr,
			logData.FileName, logData.FuncName, logData.LineNo, logData.Message)

	}
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

	logData := writeLog(LogLevelDebug, format, args...)
	select {
	case f.LogDataChan <- logData:
	//如果队列满了,走default分支,抛弃此条日志
	default:
	}

}

func (f *FileLogger) Trace(format string, args ...interface{}) {
	if f.level > LogLevelTrace {
		return
	}
	logData := writeLog(LogLevelTrace, format, args...)
	select {
	case f.LogDataChan <- logData:
	//如果队列满了,走default分支,抛弃此条日志
	default:
	}

}

func (f *FileLogger) Info(format string, args ...interface{}) {
	if f.level > LogLevelInfo {
		return
	}
	logData := writeLog(LogLevelInfo, format, args...)
	select {
	case f.LogDataChan <- logData:
	//如果队列满了,走default分支,抛弃此条日志
	default:
	}

}

func (f *FileLogger) Warn(format string, args ...interface{}) {
	if f.level > LogLevelWarn {
		return
	}
	logData := writeLog(LogLevelWarn, format, args...)
	select {
	case f.LogDataChan <- logData:
	//如果队列满了,走default分支,抛弃此条日志
	default:
	}
}

func (f *FileLogger) Error(format string, args ...interface{}) {
	if f.level > LogLevelError {
		return
	}
	logData := writeLog(LogLevelError, format, args...)
	select {
	case f.LogDataChan <- logData:
	//如果队列满了,走default分支,抛弃此条日志
	default:
	}
}

func (f *FileLogger) Fatal(format string, args ...interface{}) {
	if f.level > LogLevelFatal {
		return
	}
	logData := writeLog(LogLevelFatal, format, args...)
	select {
	case f.LogDataChan <- logData:
	//如果队列满了,走default分支,抛弃此条日志
	default:
	}
}

func (f *FileLogger) Close() {
	f.file.Close()
	f.warnFile.Close()
}
