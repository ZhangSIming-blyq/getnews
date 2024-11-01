package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// Config 配置结构体
type Config struct {
	MySQL struct {
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		DBName   string `json:"dbname"`
	} `json:"mysql"`

	Server struct {
		Port int `json:"port"`
	} `json:"server"`

	Crawler struct {
		Interval int `json:"interval"` // 爬虫间隔（分钟）
	} `json:"crawler"`
}

var (
	config *Config
	once   sync.Once
)

// Load 加载配置文件
func Load(path string) error {
	var err error
	once.Do(func() {
		var file *os.File
		file, err = os.Open(path)
		if err != nil {
			return
		}
		defer file.Close()

		config = &Config{}
		err = json.NewDecoder(file).Decode(config)
	})
	return err
}

// Get 获取配置
func Get() *Config {
	return config
}

// GetDSN 获取MySQL连接字符串
func GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.MySQL.User,
		config.MySQL.Password,
		config.MySQL.Host,
		config.MySQL.Port,
		config.MySQL.DBName,
	)
}
