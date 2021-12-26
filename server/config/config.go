package config

import (
	"io/ioutil"
	"monitor_system/global"
	"sync"

	"gopkg.in/yaml.v2"
)

type ServerConfig struct {
	Name      string `yaml:"Name"`
	RunMode   string `yaml:"RunMode"`
	AppCode   string `yaml:"AppCode"`
	HttpPort  int    `yaml:"HttpPort"`
	CipherKey string `yaml:"CipherKey"`
}

type LogConfig struct {
	LogSavePath    string `yaml:"LogSavePath"`
	LogFileName    string `yaml:"LogFileName"`
	LogFileExt     string `yaml:"LogFileExt"`
	LogFileMaxSize uint   `yaml:"LogFileMaxSize"`
}
type SessionConfig struct {
	Name   string `yaml:"Name"`
	MaxAge int    `yaml:"MaxAge"`
}

type DBConfig struct {
	DBType   string `yaml:"DBType"`
	UserName string `yaml:"UserName"`
	Password string `yaml:"Password"`
	DBName   string `yaml:"DBName"`
	Host     string `yaml:"Host"`
	Charset  string `yaml:"Charset"`
}
type Conf struct {
	Server  *ServerConfig  `yaml:"Server"`
	Log     *LogConfig     `yaml:"Log"`
	Session *SessionConfig `yaml:"Session"`
	DB      *DBConfig      `yaml:"Database"`
}

var (
	conf        *Conf
	configMutex sync.Mutex
)

func GetConfig() *Conf {
	if conf != nil {
		return conf
	}

	configMutex.Lock()
	defer configMutex.Unlock()

	// double check
	if conf != nil {
		return conf
	}

	conf = &Conf{}
	// 加载文件
	yamlFile, err := ioutil.ReadFile(global.ConfigPath)
	if err != nil {
		return nil
	}
	// 将读取的yaml文件解析为响应的 struct
	err = yaml.Unmarshal(yamlFile, conf)
	if err != nil {
		return nil
	}

	return conf
}
