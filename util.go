package logger

import (
	"fmt"
	"os"
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

func writeLog(file *os.File, level int, format string, args ...interface{}) {

	now := time.Now()
	//now.Format函数传入的时间点 必须是此时间,但是可以更改格式
	nowStr := now.Format("2006-01-02 15:04:05.999")
	levelStr := getLevelText(level)
	fileName, funcName, lineNo := GetLineInfo()
	msg := fmt.Sprintf(format, args...)
	fmt.Fprintf(file, "[%s] %s [%s:%s:%d] %s\n", nowStr, levelStr, fileName, funcName, lineNo, msg)
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
