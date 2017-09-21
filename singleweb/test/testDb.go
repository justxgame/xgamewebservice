package main

import (
	"xgamewebservice/singleweb/util"
	"xgamewebservice/singleweb/db"
	"fmt"
	"runtime/debug"
)

const configFile = "/Users/william/work/go/src/xgamewebservice/singleweb/config.json"

var serverConfig util.Config

func main() {

	// 加载资源
	// 加载配置文件
	util.LoadConfiguration(configFile, &serverConfig)
	util.LoadDbConfig(serverConfig.DBConfig)

	var serverInfo, err = db.QueryServerIpByServerId("1000")

	if nil != err {
		fmt.Printf("query failed , %s , %s ", err, debug.Stack())
	} else {
		fmt.Printf("serverInfo: %s ", serverInfo)
	}

}