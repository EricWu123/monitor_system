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

func getSessionID(sessionName string) (string, error) {
	router, err := NewRouter()
	if err != nil {
		return "", err
	}
	w := httptest.NewRecorder()
	// 登录获取session
	loginParam := make(map[string]string)
	loginParam["userName"] = "admin"
	loginParam["password"] = "BiTYKRZpk8VONUjnZ8XMeA=="
	loginRequest, _ := genReq("http://127.0.0.1:8090/login", loginParam)
	router.ServeHTTP(w, loginRequest)

	loginResp := w.Result()

	var sessionid string
	for _, cookie := range loginResp.Cookies() {
		if cookie.Name == sessionName {
			sessionid = cookie.Value
		}
	}
	return sessionid, nil
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
	type saveParam struct {
		UserName string `json:"userName"`
		Password string `json:"password"`
	}
	params := []saveParam{
		{UserName: "admin", Password: "BiTYKRZpk8VONUjnZ8XMeA=="},
		{UserName: "guest", Password: "BiTYKRZpk8VONUjnZ8XMeA=="},
		{UserName: "adaamin", Password: "BiTYKRZpk8VONUjnZ8XMeA=="},
		{UserName: "&&%", Password: "BiTYKRZpk8VONUjnZ8XMeA=="},
	}
	returnCode := []string{
		"0",
		"2",
		"2",
		"1",
	}
	for idx, value := range params {
		w := httptest.NewRecorder()
		request, _ := genReq("http://127.0.0.1:8090/login", value)
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
		assert.Equal(t, returnCode[idx], code)
	}
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

func TestSaveSystemInfo(t *testing.T) {
	config.ConfigPath = "/home/wyq/code/monitor_system/server/config/config.yaml"
	router, err := NewRouter()
	if err != nil {
		fmt.Println("post err:", err)
		t.Error("new router failed.")
		return
	}
	type saveParam struct {
		OS       string   `json:"OS"`
		HostName string   `json:"HostName"`
		IPs      []string `json:"IPs"`
	}
	params := []saveParam{
		{OS: "linux;touch test", HostName: "wyq-System-Product-Name", IPs: []string{"1.1.1.1"}},
		{OS: "linux", HostName: "*wyq-System-Product-Name", IPs: []string{"1.1.1.1"}},
		{OS: "linux", HostName: "wyq-System-Product-Name", IPs: []string{"1.11.1"}},
		{OS: "linux", HostName: "wyq-System-Product-Name", IPs: []string{"1.1.1.1"}},
	}
	returnCode := []string{
		"1",
		"1",
		"1",
		"0",
	}
	for idx, value := range params {
		w := httptest.NewRecorder()
		request, _ := genReq("http://127.0.0.1:8090/report/system_info", value)
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
		assert.Equal(t, returnCode[idx], code)
	}
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

func TestQuerySystemInfo(t *testing.T) {
	config.ConfigPath = "/home/wyq/code/monitor_system/server/config/config.yaml"
	router, err := NewRouter()
	if err != nil {
		fmt.Println("post err:", err)
		t.Error("new router failed.")
		return
	}
	sessionid, _ := getSessionID(config.GetConfig().Session.Name)

	type queryParam struct {
		OS       string `json:"OS"`
		HostName string `json:"HostName"`
		Begin    string `json:"begin"`
		End      string `json:"end"`
	}
	params := []queryParam{
		{OS: "linux;touch test", HostName: "wyq-System-Product-Name", Begin: "1639305120", End: "1639316241"},
		{OS: "linux", HostName: "*wyq-System-Product-Name", Begin: "1639305120", End: "1639316241"},
		{OS: "linux", HostName: "wyq-System-Product-Name", Begin: "1639305120", End: "bb"},
		{OS: "linux", HostName: "wyq-System-Product-Name", Begin: "aa", End: "1639316241"},
		{OS: "linux", HostName: "wyq-System-Product-Name", Begin: "1639305120", End: "1639316241"},
	}
	returnCode := []string{
		"1",
		"1",
		"1",
		"1",
		"0",
	}
	for idx, value := range params {
		w := httptest.NewRecorder()

		request, _ := genReq("http://127.0.0.1:8090/query/system_info", value)
		sessionCookies := &http.Cookie{Name: config.GetConfig().Session.Name, Value: sessionid}
		request.AddCookie(sessionCookies)
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
		assert.Equal(t, returnCode[idx], code)
	}

}
