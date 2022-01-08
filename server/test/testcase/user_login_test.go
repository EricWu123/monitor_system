package testcase

import (
	"monitor_system/test/testtools"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func TestLoginOK(t *testing.T) {
	// 前置条件

	// 执行操作
	loginParam := make(map[string]string)
	loginParam["userName"] = "admin"
	loginParam["password"] = "BiTYKRZpk8VONUjnZ8XMeA=="
	resp, err := testtools.SendHttp("http://127.0.0.1:8090/login", loginParam, nil, nil)

	//校验结果
	assert.Equal(t, nil, err)

	body, err := testtools.ReadBody(resp)
	assert.Equal(t, nil, err)

	code := gjson.Get(body, "code").String()
	assert.Equal(t, "0", code)

	//恢复环境
}

func TestLoginWrongPassword(t *testing.T) {
	// 前置条件

	// 执行操作
	loginParam := make(map[string]string)
	loginParam["userName"] = "guest"
	loginParam["password"] = "BiTYKRZpk8VONUjnZ8XMeA=="
	resp, err := testtools.SendHttp("http://127.0.0.1:8090/login", loginParam, nil, nil)

	//校验结果
	assert.Equal(t, nil, err)

	body, err := testtools.ReadBody(resp)
	assert.Equal(t, nil, err)

	code := gjson.Get(body, "code").String()
	assert.Equal(t, "2", code)

	//恢复环境
}

func TestLoginWrongUser(t *testing.T) {
	// 前置条件

	// 执行操作
	loginParam := make(map[string]string)
	loginParam["userName"] = "&&%"
	loginParam["password"] = "BiTYKRZpk8VONUjnZ8XMeA=="
	resp, err := testtools.SendHttp("http://127.0.0.1:8090/login", loginParam, nil, nil)

	//校验结果
	assert.Equal(t, nil, err)

	body, err := testtools.ReadBody(resp)
	assert.Equal(t, nil, err)

	code := gjson.Get(body, "code").String()
	assert.Equal(t, "1", code)

	//恢复环境
}
