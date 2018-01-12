package util

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	//"runtime/debug"
	"runtime/debug"
)

// 配置
type Config struct {
	Port     int      `json:"port"` // 端口号
	Host     string   `json:"host"` //host
	DBConfig DBConfig `json:"db"`   //db配置
	AppKey   string   `json:"appKey"`
	PayKey   string   `json:"payKey"`
}

type DBConfig struct {
	Address  string `json:"address"`   // instance地址
	UserName string `json:"user_name"` // 用户名
	Password string `json:"password"`  //密码
	DataBase string `json:"database"`  //连接的数据库
}

var ServerConfig Config

func LoadConfiguration(file string) {

	// 如果加载配置文件失败，则退出系统
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("load config file failed , exit system, err = %s  , %s", err, debug.Stack())
		}
	}()

	configFile, err := os.Open(file)
	defer configFile.Close()
	if nil != err {
		log.Panicf("open config file error , %s", err)
	}
	configData, err := ioutil.ReadAll(configFile)

	if nil != err {
		log.Panicf("read config file error , %s", err)
		if nil != err {
			log.Panicf("parse config file error , %s", err)
		}

		log.Printf("[config] load success config = ,\n  %v \n", string(configData))
	} else {
		err = json.Unmarshal(configData, &ServerConfig)
	}

}
