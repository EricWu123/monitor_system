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
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
)

func parseSystemInfoParam(context *gin.Context) (*model.SystemInfo, error) {
	systemInfo := model.SystemInfo{}
	e := context.ShouldBindJSON(&systemInfo)
	if e != nil {
		return nil, errors.New("bind info failed")
	}
	for _, ip := range systemInfo.IPs {
		address := net.ParseIP(ip)
		if address == nil {
			return nil, errors.New("wrong ip format")
		}
	}
	checkResult, e := utils.CheckStrWhite(systemInfo.HostName, `^[a-z-A-Z0-9-_+.]+$`, 100)
	if e != nil || !checkResult {
		return nil, errors.New("save failed, verify host name failed")
	}
	checkResult, e = utils.CheckStrWhite(systemInfo.OS, `^[a-z-A-Z0-9-_+.]+$`, 100)
	if e != nil || !checkResult {
		return nil, errors.New("save failed, verify os failed")
	}
	return &systemInfo, nil
}

func SaveSystemInfo(context *gin.Context) {
	systemInfo, e := parseSystemInfoParam(context)
	if e != nil {
		context.JSON(http.StatusOK, response.Response(errcode.ERR_CODE_INVALID_PARAMS))
		return
	}

	logging.LogInfo("system info:", systemInfo)
	conf := config.GetConfig()
	mysqlRepo := dao.NewSystemInfoMysql(conf.DB)
	e = modules.SaveSystemInfo(systemInfo, mysqlRepo)
	if e != nil {
		logging.LogInfo("save system info failed. err:", e)
		context.JSON(http.StatusOK, response.Response(errcode.ERR_CODE_FAILED))
		return
	}
	context.JSON(http.StatusOK, response.Response(errcode.ERR_CODE_OK))
}
