package logger

import (
	"fmt"
	"path"
	"runtime"
	"time"
)

func GetLineInfo() (fileName string, funcName string, lineNo int) {
	//pc 表示程序执行的计数器,Caller里的参数表示第几次调用函数的地方
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		//对返回值取绝对路径
		fileName = path.Base(file)
		funcName = path.Base(runtime.FuncForPC(pc).Name())
		lineNo = line
	}
	return
}

/*
1.当业务调用打日志的方法时,我们把日志相关的数据写入chan队列
2.然后后台启动一个线程,不短的从chan里获取这些日志,写入到文件里
*/

type LogData struct {
	Message      string
	TimeStr      string
	LevelStr     string
	FileName     string
	FuncName     string
	LineNo       int
	WarnAndFatal bool
}

func writeLog(level int, format string, args ...interface{}) *LogData {

	now := time.Now()
	//now.Format函数传入的时间点 必须是此时间,但是可以更改格式
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := getLevelText(level)
	fileName, funcName, lineNo := GetLineInfo()
	msg := fmt.Sprintf(format, args...)
	logData := &LogData{
		Message:      msg,
		TimeStr:      nowStr,
		LevelStr:     levelStr,
		FileName:     fileName,
		FuncName:     funcName,
		LineNo:       lineNo,
		WarnAndFatal: false,
	}
	if level == LogLevelError || level == LogLevelWarn || level == LogLevelFatal {
		logData.WarnAndFatal = true
	}
	return logData
	//fmt.Fprintf(file, "[%s] %s [%s:%s:%d] %s\n", nowStr, levelStr, fileName, funcName, lineNo, msg)
}

func getLevelText(level int) string {
	switch level {
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelTrace:
		return "TRACE"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	case LogLevelFatal:
		return "FATAL"
	default:
		return "UNKNOW"

	}
}

func getLogLevel(level string) int {
	switch level {
	case "debug":
		return LogLevelDebug
	case "trace":
		return LogLevelTrace
	case "info":
		return LogLevelTrace
	case "warn":
		return LogLevelWarn
	case "error":
		return LogLevelError
	case "fatal":
		return LogLevelFatal
	default:
		return LogLevelDebug

	}
}
