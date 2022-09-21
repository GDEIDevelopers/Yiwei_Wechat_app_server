package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	// HTTP服务器监听地址端口
	// 默认为0.0.0.0:5607
	HTTPAddr string `json:"httpaddr, omitempty"`
	// MySQL数据库地址端口
	// 默认为 127.0.0.1:3306
	DBAddr string `json:"dbaddr, omitempty"`
	// MySQL数据库用户
	DBUser string `json:"dbuser"`
	// MySQL数据库密码
	DBPass string `json:"dbpass"`
	// MySQL数据库数据库
	DBBase string `json:"dbbase"`
	// SSL证书(选填)，不填不使用HTTPS
	Cert string `json:"certfile, omitempty"`
	// SSL证书密钥，不填不使用HTTPS
	Key string `json:"keyfile, omitempty"`
}

func Read(filename string) *Config {
	_file, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	config := &Config{}
	if err := json.Unmarshal(_file, config); err != nil {
		log.Fatal(err)
	}

	if config.HTTPAddr == "" {
		config.HTTPAddr = "0.0.0.0:5607"
	}

	if config.DBAddr == "" {
		config.DBAddr = "127.0.0.1:3306"
	}

	if config.DBBase == "" {
		config.DBBase = "net"
	}
	return config
}
