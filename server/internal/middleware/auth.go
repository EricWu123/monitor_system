package middleware

import (
	"monitor_system/config"
	"monitor_system/global"
	"monitor_system/internal/dao"
	"monitor_system/logging"
	"monitor_system/modules"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Autheticate(context *gin.Context) {
	conf := config.GetConfig()
	sess := dao.NewMysqlSession(conf)
	sessID, err := context.Request.Cookie("sessionid")
	if err != nil {
		logging.LogInfo("no rights access. err:", err)
		context.JSON(http.StatusOK, gin.H{"code": global.FAILED, "msg": "no rights access."})
		context.Abort()
		return
	}
	ret, err := modules.IsValid(sessID.Value, sess)
	if err != nil || !ret {
		logging.LogInfo("no rights access. err:", err)
		context.JSON(http.StatusOK, gin.H{"code": global.FAILED, "msg": "no rights access."})
		context.Abort()
		return
	}
	modules.Update(sessID.Value, sess)
	logging.LogInfo("auth success. sess:", sessID.Value)
}

func ReportAuth(context *gin.Context) {
	conf := config.GetConfig()
	authorization := context.Request.Header.Get("Authorization")
	if authorization != "APPCODE "+conf.Server.AppCode {
		logging.LogInfo("appcode is not valid.")
		context.JSON(http.StatusOK, gin.H{"code": global.FAILED, "msg": "appcode is not valid."})
		context.Abort()
	}
	logging.LogInfo("ReportAuth success.")
}
