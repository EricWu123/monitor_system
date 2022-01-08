package testcase

import (
	"monitor_system/config"
	"monitor_system/test/testtools"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestSaveSystemInfoOK(t *testing.T) {
	// 前置条件

	// 执行操作
	param := make(map[string]interface{})
	param["HostName"] = "wyq"
	param["OS"] = "Linux"
	param["IPs"] = []string{"1.1.1.1"}

	header := make(map[string]string)
	header["Authorization"] = "APPCODE " + config.GetConfig().Server.AppCode

	resp, err := testtools.SendHttp("http://127.0.0.1:8090/report/system_info", param, header, nil)

	//校验结果
	assert.Equal(t, nil, err)

	body, err := testtools.ReadBody(resp)
	assert.Equal(t, nil, err)

	code := gjson.Get(body, "code").String()
	assert.Equal(t, "0", code)

	//恢复环境
}

func TestSaveSystemInfoWrongOS(t *testing.T) {
	// 前置条件

	// 执行操作
	param := make(map[string]interface{})
	param["HostName"] = "wyq"
	param["OS"] = "Lin;ux"
	param["IPs"] = []string{"1.1.1.1"}

	header := make(map[string]string)
	header["Authorization"] = "APPCODE " + config.GetConfig().Server.AppCode

	resp, err := testtools.SendHttp("http://127.0.0.1:8090/report/system_info", param, header, nil)

	//校验结果
	assert.Equal(t, nil, err)

	body, err := testtools.ReadBody(resp)
	assert.Equal(t, nil, err)

	code := gjson.Get(body, "code").String()
	assert.Equal(t, "1", code)

	//恢复环境
}

func TestSaveSystemInfoWrongIP(t *testing.T) {
	// 前置条件

	// 执行操作
	param := make(map[string]interface{})
	param["HostName"] = "wyq"
	param["OS"] = "Linux"
	param["IPs"] = []string{"1.11.1"}

	header := make(map[string]string)
	header["Authorization"] = "APPCODE " + config.GetConfig().Server.AppCode

	resp, err := testtools.SendHttp("http://127.0.0.1:8090/report/system_info", param, header, nil)

	//校验结果
	assert.Equal(t, nil, err)

	body, err := testtools.ReadBody(resp)
	assert.Equal(t, nil, err)

	code := gjson.Get(body, "code").String()
	assert.Equal(t, "1", code)

	//恢复环境
}

func TestSaveSystemInfoWrongHostName(t *testing.T) {
	// 前置条件

	// 执行操作
	param := make(map[string]interface{})
	param["HostName"] = "*sldkfj"
	param["OS"] = "Linux"
	param["IPs"] = []string{"1.11.1"}

	header := make(map[string]string)
	header["Authorization"] = "APPCODE " + config.GetConfig().Server.AppCode

	resp, err := testtools.SendHttp("http://127.0.0.1:8090/report/system_info", param, header, nil)

	//校验结果
	assert.Equal(t, nil, err)

	body, err := testtools.ReadBody(resp)
	assert.Equal(t, nil, err)

	code := gjson.Get(body, "code").String()
	assert.Equal(t, "1", code)

	//恢复环境
}
