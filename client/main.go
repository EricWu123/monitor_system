package main

import (
	"errors"
	"fmt"
	"log"
	"monitor_system_client/config"
	"os"
	"time"
)

var gConf *config.Conf
var gLog *log.Logger

func initConf() error {
	gConf = new(config.Conf)
	if gConf == nil {
		return errors.New("conf new failed")
	}
	return gConf.InitConfig()
}

func initLog() error {
	fileName := gConf.Log.LogSavePath + gConf.Log.LogFileName + gConf.Log.LogFileExt
	logFile, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	gLog = log.New(logFile, "["+gConf.Server.Name+"] ", log.Ldate|log.Ltime|log.Lshortfile)
	if gLog == nil {
		return errors.New("log new failed")
	}
	return nil
}

func main() {
	err := initConf()
	if err != nil {
		fmt.Println("init conf failed. err:", err)
		return
	}
	err = initLog()
	if err != nil {
		fmt.Println("init log failed. err:", err)
		return
	}
	c := time.Tick(time.Duration(gConf.Server.Interval) * time.Second)
	for {
		s := systemInfo{}
		s.getHostInfo()
		s.getNetInfo()
		gLog.Printf("upload info:%v", s)
		err = s.upload("http://" + gConf.Server.Addr + "/report/system_info")
		if err != nil {
			gLog.Printf("upload info error:%v", err)
		} else {
			gLog.Printf("upload info success")
		}
		<-c
	}
}
