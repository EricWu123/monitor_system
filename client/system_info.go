package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"

	"github.com/shirou/gopsutil/host"
	"github.com/tidwall/gjson"
)

type systemInfo struct {
	HostName string   `json:"HostName"`
	OS       string   `json:"OS"`
	IPs      []string `json:"IPs"`
}

func (s *systemInfo) getHostInfo() error {
	n, err := host.Info()
	if err != nil {
		return err
	}
	s.HostName = n.Hostname
	s.OS = n.OS
	return nil
}
func getLocalIPs() (ips []string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		ips = append(ips, ipAddr.IP.String())
	}
	return
}

func (s *systemInfo) getNetInfo() {
	ips, _ := getLocalIPs()
	s.IPs = ips
}

func (s *systemInfo) upload(url string) error {
	client := &http.Client{}

	infoStr, err := json.Marshal(s)
	if err != nil {
		return err
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(infoStr))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "APPCODE 7bb68fc49491c180ce6002f99d06db7f")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	code := gjson.Get(string(body), "code").Int()
	if code != 0 {
		return errors.New(string(body))
	}
	return nil
}
