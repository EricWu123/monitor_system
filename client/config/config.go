package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Name     string `yaml:"Name"`
	RunMode  string `yaml:"RunMode"`
	AppCode  string `yaml:"AppCode"`
	Addr     string `yaml:"Addr"`
	Interval int    `yaml:"Interval"`
}

type LogConfig struct {
	LogSavePath    string `yaml:"LogSavePath"`
	LogFileName    string `yaml:"LogFileName"`
	LogFileExt     string `yaml:"LogFileExt"`
	LogFileMaxSize uint   `yaml:"LogFileMaxSize"`
}

type Conf struct {
	Server *ServerConfig `yaml:"Server"`
	Log    *LogConfig    `yaml:"Log"`
}

func (c *Conf) InitConfig() error {
	yamlFile, err := ioutil.ReadFile("./config/config.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return err
	}
	return nil
}
