package model

type SystemInfo struct {
	HostName string   `json:"HostName"`
	OS       string   `json:"OS"`
	IPs      []string `json:"IPs"`
}
