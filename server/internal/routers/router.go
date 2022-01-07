package router

import (
	"errors"
	"monitor_system/apis"
	"monitor_system/config"
	"monitor_system/internal/db"
	"monitor_system/internal/middleware"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
)

func NewRouter() (*gin.Engine, error) {
	conf := config.GetConfig()
	db := db.NewMysql(conf.DB)
	if db == nil {
		return nil, errors.New("init db failed")
	}

	gin.SetMode(conf.Server.RunMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/report/system_info", middleware.ReportAuth, apis.SaveSystemInfo)
	r.POST("/query/system_info", middleware.Autheticate, apis.QuerySystemInfo)
	r.POST("/login", apis.UserLogin)
	r.POST("/", apis.UserLogin)

	if conf.Server.RunMode == "debug" {
		pprof.Register(r)
	}

	return r, nil
}
