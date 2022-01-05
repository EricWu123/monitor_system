package router

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"monitor_system/config"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func genReq(url string, param interface{}) (*http.Request, error) {
	infoStr, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(infoStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func readBody(r *http.Response) (string, error) {
	if r == nil {
		return "", errors.New("nil resp")
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	fmt.Printf("resp body:%+v\n", string(body))
	return string(body), nil
}

func TestLogin(t *testing.T) {
	config.ConfigPath = "/home/wyq/code/monitor_system/server/config/config.yaml"
	router, err := NewRouter()
	if err != nil {
		fmt.Println("post err:", err)
		t.Error("new router failed.")
		return
	}
	w := httptest.NewRecorder()
	param := make(map[string]string)
	param["userName"] = "admin"
	param["password"] = "BiTYKRZpk8VONUjnZ8XMeA=="
	request, _ := genReq("http://127.0.0.1:8090/login", param)
	router.ServeHTTP(w, request)

	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)

	body, err := readBody(resp)
	if err != nil {
		fmt.Println("post err:", err)
		t.Error("post failed.")
		return
	}
	fmt.Println("body:", body)
	code := gjson.Get(body, "code").String()
	assert.Equal(t, "0", code)
}

func TestSaveSystemInfoNorights(t *testing.T) {
	config.ConfigPath = "/home/wyq/code/monitor_system/server/config/config.yaml"
	router, err := NewRouter()
	if err != nil {
		fmt.Println("post err:", err)
		t.Error("new router failed.")
		return
	}
	w := httptest.NewRecorder()
	param := make(map[string]interface{})
	param["HostName"] = "wyq"
	param["OS"] = "Linux"
	param["IPs"] = []string{"1.1.1.1"}
	request, _ := genReq("http://127.0.0.1:8090/report/system_info", param)
	router.ServeHTTP(w, request)

	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)

	body, err := readBody(resp)
	if err != nil {
		fmt.Println("post err:", err)
		t.Error("post failed.")
		return
	}
	fmt.Println("body:", body)
	code := gjson.Get(body, "code").String()
	assert.Equal(t, "3", code)
}

func TestSaveSystemInfoOK(t *testing.T) {
	config.ConfigPath = "/home/wyq/code/monitor_system/server/config/config.yaml"
	router, err := NewRouter()
	if err != nil {
		fmt.Println("post err:", err)
		t.Error("new router failed.")
		return
	}
	w := httptest.NewRecorder()
	param := make(map[string]interface{})
	param["HostName"] = "wyq-System-Product-Name"
	param["OS"] = "linux"
	param["IPs"] = []string{"1.1.1.1"}
	request, _ := genReq("http://127.0.0.1:8090/report/system_info", param)
	request.Header.Add("Authorization", "APPCODE "+config.GetConfig().Server.AppCode)
	router.ServeHTTP(w, request)

	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)

	body, err := readBody(resp)
	if err != nil {
		fmt.Println("post err:", err)
		t.Error("post failed.")
		return
	}
	fmt.Println("body:", body)
	code := gjson.Get(body, "code").String()
	assert.Equal(t, "0", code)
}

func TestQuerySystemInfoNoRights(t *testing.T) {
	config.ConfigPath = "/home/wyq/code/monitor_system/server/config/config.yaml"
	router, err := NewRouter()
	if err != nil {
		fmt.Println("post err:", err)
		t.Error("new router failed.")
		return
	}
	w := httptest.NewRecorder()
	param := make(map[string]interface{})
	param["begin"] = "1639305120"
	param["end"] = "1639316241"
	param["HostName"] = "wyq-System-Product-Name"
	param["OS"] = "linux"
	request, _ := genReq("http://127.0.0.1:8090/query/system_info", param)
	request.Header.Add("Authorization", "APPCODE "+config.GetConfig().Server.AppCode)
	router.ServeHTTP(w, request)

	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)

	body, err := readBody(resp)
	if err != nil {
		fmt.Println("post err:", err)
		t.Error("post failed.")
		return
	}
	fmt.Println("body:", body)
	code := gjson.Get(body, "code").String()
	assert.Equal(t, "3", code)
}

func TestQuerySystemInfoOK(t *testing.T) {
	config.ConfigPath = "/home/wyq/code/monitor_system/server/config/config.yaml"
	router, err := NewRouter()
	if err != nil {
		fmt.Println("post err:", err)
		t.Error("new router failed.")
		return
	}
	w := httptest.NewRecorder()
	// 登录获取session
	loginParam := make(map[string]string)
	loginParam["userName"] = "admin"
	loginParam["password"] = "BiTYKRZpk8VONUjnZ8XMeA=="
	request, _ := genReq("http://127.0.0.1:8090/login", loginParam)
	router.ServeHTTP(w, request)

	resp := w.Result()
	assert.Equal(t, 200, resp.StatusCode)

	var sessionid string
	for _, cookie := range resp.Cookies() {
		if cookie.Name == config.GetConfig().Session.Name {
			sessionid = cookie.Value
		}
	}

	param := make(map[string]interface{})
	param["begin"] = "1639305120"
	param["end"] = "1639316241"
	param["HostName"] = "wyq-System-Product-Name"
	param["OS"] = "linux"
	request, _ = genReq("http://127.0.0.1:8090/query/system_info", param)
	sessionCookies := &http.Cookie{Name: config.GetConfig().Session.Name, Value: sessionid}
	request.AddCookie(sessionCookies)
	router.ServeHTTP(w, request)

	resp = w.Result()
	assert.Equal(t, 200, resp.StatusCode)

	body, err := readBody(resp)
	if err != nil {
		fmt.Println("post err:", err)
		t.Error("post failed.")
		return
	}
	fmt.Println("body:", body)
	code := gjson.Get(body, "code").String()
	assert.Equal(t, "0", code)
}
