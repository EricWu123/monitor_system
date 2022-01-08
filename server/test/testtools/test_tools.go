package testtools

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"monitor_system/config"
	router "monitor_system/internal/routers"
	"monitor_system/logging"
	"net/http"
	"strconv"
)

func TestRunServer() {
	go func() {
		config.ConfigPath = "../../config/config.yaml"
		conf := config.GetConfig()
		if conf == nil {
			return
		}
		logging.LogInfo("app start...")
		router, err := router.NewRouter()
		if err != nil {
			logging.LogInfo("new router failed. err:", err)
			return
		}

		router.Run(":" + strconv.Itoa(conf.Server.HttpPort))
	}()

}

func ShutdownServer() {

}

func SendHttp(url string, param interface{}, header map[string]string, c *http.Cookie) (*http.Response, error) {
	infoStr, err := json.Marshal(param)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(infoStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	for key, value := range header {
		req.Header.Set(key, value)
	}
	if c != nil {
		req.AddCookie(c)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	return resp, err
}

func ReadBody(r *http.Response) (string, error) {
	if r == nil {
		return "", errors.New("nil resp")
	}
	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetSessionID(sessionName string) (string, error) {

	loginParam := make(map[string]string)
	loginParam["userName"] = "admin"
	loginParam["password"] = "BiTYKRZpk8VONUjnZ8XMeA=="
	resp, err := SendHttp("http://127.0.0.1:8090/login", loginParam, nil, nil)

	if err != nil {
		return "", err
	}
	var sessionid string
	for _, cookie := range resp.Cookies() {
		if cookie.Name == sessionName {
			sessionid = cookie.Value
		}
	}
	return sessionid, nil
}
