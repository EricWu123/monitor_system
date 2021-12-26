package router

import (
	"errors"
	"monitor_system/apis"
	"monitor_system/config"
	"monitor_system/internal/db"
	"monitor_system/internal/middleware"
	"monitor_system/session"

	"github.com/gin-gonic/gin"
)

func NewRouter() (*gin.Engine, error) {
	conf := config.GetConfig()
	db := db.NewMysql(conf.DB)
	if db == nil {
		return nil, errors.New("init db failed")
	}
	session := session.NewSession(db)
	auth := middleware.NewAuth(session, conf)

	gin.SetMode(conf.Server.RunMode)
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.POST("/report/system_info", auth.ReportAuth, apis.SaveSystemInfo)
	r.POST("/query/system_info", auth.Autheticate, apis.SystemInfo)
	r.POST("/login", apis.UserLogin)
	r.POST("/", apis.UserLogin)

	return r, nil
}
