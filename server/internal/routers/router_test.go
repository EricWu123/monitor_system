package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

func post(url string, param interface{}) (string, error) {
	client := &http.Client{}

	infoStr, err := json.Marshal(param)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(infoStr))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func TestLogin(t *testing.T) {
	resp, err := post("http://127.0.0.1:8090/login", nil)
	if err != nil {
		fmt.Println("post err:", err)
	}
	code := gjson.Get(resp, "code").String()
	assert.Equal(t, "1", code)

	param := make(map[string]string)
	param["userName"] = "admin"
	param["password"] = "BiTYKRZpk8VONUjnZ8XMeA=="
	resp, err = post("http://127.0.0.1:8090/login", param)
	if err != nil {
		fmt.Println("post err:", err)
	}
	code = gjson.Get(resp, "code").String()
	assert.Equal(t, "0", code)
}
