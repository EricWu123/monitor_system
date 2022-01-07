package main

import (
	"flag"
	"monitor_system/config"
	router "monitor_system/internal/routers"
	"monitor_system/logging"
	"strconv"
)

func main() {
	configPath := flag.String("configPath", "./config/config.yaml", "配置文件路径")
	config.ConfigPath = *configPath
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

	router.Run(":" + strconv.Itoa(conf.Server.HttpPort))
}
