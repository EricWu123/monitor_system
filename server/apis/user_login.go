package apis

import (
	"errors"
	"monitor_system/config"
	"monitor_system/errcode"
	"monitor_system/internal/dao"
	"monitor_system/internal/utils"
	"monitor_system/logging"
	"monitor_system/model"
	"monitor_system/modules"
	"monitor_system/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func parseLoginParam(context *gin.Context) (map[string]string, error) {
	param := make(map[string]string)
	e := context.BindJSON(&param)
	if e != nil {
		return nil, errors.New("bind failed")
	}
	checkResult, e := utils.CheckStrWhite(param["userName"], `^[a-z-A-Z0-9]+$`, 100)
	if e != nil || !checkResult {
		return nil, errors.New("login failed, verify failed")
	}
	return param, nil
}

func UserLogin(context *gin.Context) {
	param, e := parseLoginParam(context)
	if e != nil {
		context.JSON(http.StatusOK, response.Response(errcode.ERR_CODE_INVALID_PARAMS))
		return
	}
	var user model.User
	var authResult bool
	user.UserName = param["userName"]
	conf := config.GetConfig()
	mysqlRepo := dao.NewUserMysql(conf)
	if authResult, e = modules.LoginAuth(param["password"], param["userName"], mysqlRepo); e != nil {
		logging.LogInfo("login failed. user name:", user.UserName, ",err:", e)
		context.JSON(http.StatusOK, response.Response(errcode.ERR_CODE_FAILED))
		return
	}
	if !authResult {
		logging.LogInfo("login failed. user name:", user.UserName)
		context.JSON(http.StatusOK, response.Response(errcode.ERR_CODE_FAILED))
		return
	}

	sess := dao.NewMysqlSession(conf)

	sessionID, e := modules.Set(user.UserName, sess)
	if e != nil {
		context.JSON(http.StatusOK, response.Response(errcode.ERR_CODE_FAILED))
		return
	}
	context.SetCookie(conf.Session.Name, sessionID, conf.Session.MaxAge, "/", "127.0.0.1", false, true)

	context.JSON(http.StatusOK, response.Response(errcode.ERR_CODE_OK))

}
