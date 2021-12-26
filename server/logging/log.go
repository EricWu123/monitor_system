package logging

import (
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
	fileName := conf.Log.LogSavePath + conf.Log.LogFileName + conf.Log.LogFileExt
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
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
	mylog.Println(v...)
}
