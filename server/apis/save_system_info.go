package apis

import (
	"errors"
	"monitor_system/global"
	"monitor_system/internal/model"
	"monitor_system/internal/utils"
	"monitor_system/logging"
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
		context.JSON(http.StatusOK, gin.H{"code": global.FAILED, "msg": e})
		return
	}

	logging.LogInfo("system info:%v", systemInfo)
	e = systemInfo.Save()
	if e != nil {
		logging.LogInfo("save system info failed. err:", e)
		context.JSON(http.StatusOK, gin.H{"code": global.FAILED, "msg": "save info failed"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"code": global.SUCCESS, "msg": "save system info success"})
}
