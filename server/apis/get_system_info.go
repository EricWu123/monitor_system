package apis

import (
	"errors"
	"monitor_system/config"
	"monitor_system/errcode"
	"monitor_system/internal/dao"
	"monitor_system/internal/utils"
	"monitor_system/logging"
	"monitor_system/modules"
	"monitor_system/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

func parseQueryParam(context *gin.Context) (string, error) {
	param := make(map[string]string)
	e := context.BindJSON(&param)
	if e != nil {
		return "", errors.New("bind failed")
	}
	OS := param["OS"]                // 操作系统类型
	hostName := param["HostName"]    // 主机名
	timeRangeBegin := param["begin"] // 时间范围
	timeRangeEnd := param["end"]     // 时间范围

	result, e := utils.CheckStrWhite(OS, `^[0-9a-zA-Z]+$`, 100)
	if e != nil || !result {
		return "", errors.New("get system info failed, verify OS failed")
	}
	result, e = utils.CheckStrBlack(hostName, `[!@#$%^&*]+`, 100)
	if e != nil || !result {
		return "", errors.New("get system info failed, verify hostname failed")
	}
	result, e = utils.CheckStrWhite(timeRangeEnd, `^[0-9]+$`, -1)
	if e != nil || !result {
		return "", errors.New("get system info failed, verify end failed")
	}

	result, e = utils.CheckStrWhite(timeRangeBegin, `^[0-9]+$`, -1)
	if e != nil || !result {
		return "", errors.New("get system info failed, verify start failed")
	}

	if timeRangeEnd == "" || timeRangeBegin == "" {
		return "", errors.New("get system info failed, start or end is nil")
	}

	condition := " where "
	if OS != "" {
		condition = condition + "OS = '" + OS + "' and "
	}

	if hostName != "" {
		condition = condition + "hostName = '" + hostName + "' and "
	}

	if timeRangeBegin != "" && timeRangeEnd != "" {
		condition = condition + "created <= " + timeRangeEnd + " and created >= " + timeRangeBegin
	}
	return condition, nil
}

func QuerySystemInfo(context *gin.Context) {
	var queryCondition string
	var err error
	if queryCondition, err = parseQueryParam(context); err != nil {
		logging.LogInfo("get system info failed. err:", err)
		context.JSON(http.StatusOK, response.Response(errcode.ERR_CODE_INVALID_PARAMS))
		return
	}
	conf := config.GetConfig()
	mysqlRepo := dao.NewSystemInfoMysql(conf.DB)
	sysInfos, err := modules.GetSystemInfo(queryCondition, mysqlRepo)
	if err != nil {
		logging.LogInfo("get system info failed. err:", err)
		context.JSON(http.StatusOK, response.Response(errcode.ERR_CODE_FAILED))
		return
	}
	logging.LogInfof("get system info success, info:%+v", sysInfos)

	context.JSON(http.StatusOK, response.ResponseWithData(errcode.ERR_CODE_OK, sysInfos))
}
