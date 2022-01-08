package testcase

import (
	"monitor_system/config"
	"monitor_system/test/testtools"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestQuerySystemInfoOK(t *testing.T) {
	// 前置条件

	// 执行操作
	param := make(map[string]interface{})
	param["begin"] = "1639305120"
	param["end"] = "1639316241"
	param["HostName"] = "wyq-System-Product-Name"
	param["OS"] = "linux"

	sessionid, err := testtools.GetSessionID(config.GetConfig().Session.Name)
	assert.Equal(t, nil, err)

	sessionCookies := &http.Cookie{Name: config.GetConfig().Session.Name, Value: sessionid}

	resp, err := testtools.SendHttp("http://127.0.0.1:8090/query/system_info", param, nil, sessionCookies)

	//校验结果
	assert.Equal(t, nil, err)

	body, err := testtools.ReadBody(resp)
	assert.Equal(t, nil, err)

	code := gjson.Get(body, "code").String()
	assert.Equal(t, "0", code)

	//恢复环境
}

func TestQuerySystemInfoWrongBeginParam(t *testing.T) {
	// 前置条件

	// 执行操作
	param := make(map[string]interface{})
	param["begin"] = "aa"
	param["end"] = "1639316241"
	param["HostName"] = "wyq-System-Product-Name"
	param["OS"] = "linux"

	sessionid, err := testtools.GetSessionID(config.GetConfig().Session.Name)
	assert.Equal(t, nil, err)

	sessionCookies := &http.Cookie{Name: config.GetConfig().Session.Name, Value: sessionid}

	resp, err := testtools.SendHttp("http://127.0.0.1:8090/query/system_info", param, nil, sessionCookies)

	//校验结果
	assert.Equal(t, nil, err)

	body, err := testtools.ReadBody(resp)
	assert.Equal(t, nil, err)

	code := gjson.Get(body, "code").String()
	assert.Equal(t, "1", code)

	//恢复环境
}

func TestQuerySystemInfoWrongEndParam(t *testing.T) {
	// 前置条件

	// 执行操作
	param := make(map[string]interface{})
	param["begin"] = "1639305120"
	param["end"] = "bb"
	param["HostName"] = "wyq-System-Product-Name"
	param["OS"] = "linux"

	sessionid, err := testtools.GetSessionID(config.GetConfig().Session.Name)
	assert.Equal(t, nil, err)

	sessionCookies := &http.Cookie{Name: config.GetConfig().Session.Name, Value: sessionid}

	resp, err := testtools.SendHttp("http://127.0.0.1:8090/query/system_info", param, nil, sessionCookies)

	//校验结果
	assert.Equal(t, nil, err)

	body, err := testtools.ReadBody(resp)
	assert.Equal(t, nil, err)

	code := gjson.Get(body, "code").String()
	assert.Equal(t, "1", code)

	//恢复环境
}

func TestQuerySystemInfoWrongHostNameParam(t *testing.T) {
	// 前置条件

	// 执行操作
	param := make(map[string]interface{})
	param["begin"] = "1639305120"
	param["end"] = "1639316241"
	param["HostName"] = "*wyq-System-Product-Name"
	param["OS"] = "linux"

	sessionid, err := testtools.GetSessionID(config.GetConfig().Session.Name)
	assert.Equal(t, nil, err)

	sessionCookies := &http.Cookie{Name: config.GetConfig().Session.Name, Value: sessionid}

	resp, err := testtools.SendHttp("http://127.0.0.1:8090/query/system_info", param, nil, sessionCookies)

	//校验结果
	assert.Equal(t, nil, err)

	body, err := testtools.ReadBody(resp)
	assert.Equal(t, nil, err)

	code := gjson.Get(body, "code").String()
	assert.Equal(t, "1", code)

	//恢复环境
}

func TestQuerySystemInfoWrongOSParam(t *testing.T) {
	// 前置条件

	// 执行操作
	param := make(map[string]interface{})
	param["begin"] = "1639305120"
	param["end"] = "1639316241"
	param["HostName"] = "wyq-System-Product-Name"
	param["OS"] = "li;nux"

	sessionid, err := testtools.GetSessionID(config.GetConfig().Session.Name)
	assert.Equal(t, nil, err)

	sessionCookies := &http.Cookie{Name: config.GetConfig().Session.Name, Value: sessionid}

	resp, err := testtools.SendHttp("http://127.0.0.1:8090/query/system_info", param, nil, sessionCookies)

	//校验结果
	assert.Equal(t, nil, err)

	body, err := testtools.ReadBody(resp)
	assert.Equal(t, nil, err)

	code := gjson.Get(body, "code").String()
	assert.Equal(t, "1", code)

	//恢复环境
}

func TestQuerySystemInfoNoRights(t *testing.T) {
	// 前置条件

	// 执行操作
	param := make(map[string]interface{})
	param["begin"] = "1639305120"
	param["end"] = "1639316241"
	param["HostName"] = "wyq-System-Product-Name"
	param["OS"] = "linux"

	resp, err := testtools.SendHttp("http://127.0.0.1:8090/query/system_info", param, nil, nil)

	//校验结果
	assert.Equal(t, nil, err)

	body, err := testtools.ReadBody(resp)
	assert.Equal(t, nil, err)

	code := gjson.Get(body, "code").String()
	assert.Equal(t, "3", code)

	//恢复环境
}
