package main

import (
	"monitor_system/config"
	"monitor_system/errcode"
	router "monitor_system/internal/routers"
	"monitor_system/logging"
	"strconv"
)

func main() {
	conf := config.GetConfig()
	if conf == nil {
		return
	}
	logging.LogInfo("app start...")
	router, err := router.NewRouter()
	if err != nil {
		logging.LogInfo("new router failed. err:", err)
		return
	}
	logging.LogInfo(errcode.ERR_CODE_FAILED)
	router.Run(":" + strconv.Itoa(conf.Server.HttpPort))
}
