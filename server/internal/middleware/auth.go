package middleware

import (
	"monitor_system/config"
	"monitor_system/errcode"
	"monitor_system/internal/dao"
	"monitor_system/logging"
	"monitor_system/modules"
	"monitor_system/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Autheticate(context *gin.Context) {
	conf := config.GetConfig()
	sess := dao.NewMysqlSession(conf)
	sessID, err := context.Request.Cookie("sessionid")
	if err != nil {
		logging.LogInfo("no rights access. err:", err)
		context.JSON(http.StatusOK, response.Response(errcode.ERR_CODE_NO_RIGHTS))
		context.Abort()
		return
	}
	ret, err := modules.IsValid(sessID.Value, sess)
	if err != nil || !ret {
		logging.LogInfo("no rights access. err:", err)
		context.JSON(http.StatusOK, response.Response(errcode.ERR_CODE_NO_RIGHTS))
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
		logging.LogInfo("appcode is not valid. authorization:", authorization, conf.Server.AppCode)
		context.JSON(http.StatusOK, response.Response(errcode.ERR_CODE_NO_RIGHTS))
		context.Abort()
		return
	}
	logging.LogInfo("ReportAuth success.")
}
