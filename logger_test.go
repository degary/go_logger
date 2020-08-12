package logger

import (
	"testing"
	"time"
)

func initLogger(name string, logPath, logName string, level string) {
	config := make(map[string]string, 8)
	config["log_path"] = logPath
	config["log_name"] = logName
	config["log_level"] = level
	config["log_split_type"] = "Hour"
	err := InitLogger(name, config)
	//log = logger.NewConsoleLogger(level)
	if err != nil {
		return
	}
	Debug("init logger success")
	return
}
func Run() {
	for {
		Debug("user server is running,test the func of split base on size")
		time.Sleep(1 * time.Second)

	}
}

func TestLogger(t *testing.T) {
	initLogger("file", "F:\\log\\", "user_server", "debug")
	Run()
}
