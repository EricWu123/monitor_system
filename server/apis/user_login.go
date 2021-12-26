package apis

import (
	"errors"
	"monitor_system/global"
	"monitor_system/internal/model"
	"monitor_system/internal/utils"
	"monitor_system/logging"
	"net/http"

	"github.com/gin-gonic/gin"
)

func parseLoginParam(context *gin.Context) (map[string]string, error) {
	param := make(map[string]string)
	e := context.BindJSON(&param)
	if e != nil {
		context.JSON(http.StatusOK, gin.H{"code": global.FAILED, "msg": "failed"})
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
		context.JSON(http.StatusOK, gin.H{"code": global.FAILED, "msg": e})
		return
	}
	var user model.User
	var authResult bool
	user.UserName = param["userName"]
	if authResult, e = user.LoginAuth(param["password"]); e != nil {
		logging.LogInfo("login failed. user name:", user.UserName, ",err:", e)
		context.JSON(http.StatusOK, gin.H{"code": global.FAILED, "msg": "login failed. please try again."})
		return
	}
	if !authResult {
		logging.LogInfo("login failed. user name:", user.UserName)
		context.JSON(http.StatusOK, gin.H{"code": global.FAILED, "msg": "login failed. please try again."})
		return
	}

	e = user.SetSession(context)
	if e != nil {
		context.JSON(http.StatusOK, gin.H{"code": global.FAILED, "msg": "login failed. set session failed."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"code": global.SUCCESS, "msg": "login success"})

}
