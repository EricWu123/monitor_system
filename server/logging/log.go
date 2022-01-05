package logging

import (
	"fmt"
	"log"
	"monitor_system/config"
	"os"
	"sync"
)

var (
	mylog    *log.Logger
	logMutex sync.Mutex
)

func initLog() *log.Logger {
	if mylog != nil {
		return mylog
	}

	logMutex.Lock()
	defer logMutex.Unlock()

	// double check
	if mylog != nil {
		return mylog
	}

	conf := config.GetConfig()

	os.MkdirAll(conf.Log.LogSavePath, 0775)
	fileName := conf.Log.LogSavePath + conf.Log.LogFileName + conf.Log.LogFileExt
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0664)
	if err != nil {
		return nil
	}
	mylog = log.New(logFile, "["+conf.Server.Name+"] ", log.Ldate|log.Ltime|log.Lshortfile)
	return mylog
}

func LogInfo(v ...interface{}) {
	if mylog == nil {
		mylog = initLog()
	}
	if mylog == nil {
		fmt.Println(v...)
		return
	}
	mylog.Println(v...)
}

func LogInfof(format string, a ...interface{}) {
	if mylog == nil {
		mylog = initLog()
	}
	if mylog == nil {
		fmt.Printf(format, a...)
		return
	}
	mylog.Printf(format, a...)
}
