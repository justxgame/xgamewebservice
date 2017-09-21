package main

import (
	"net/http"
	"strconv"
	"log"
	"xgamewebservice/singleweb/handler"
	_ "xgamewebservice/singleweb/util"
	"xgamewebservice/singleweb/util"
	"os"
	"fmt"
	"runtime/debug"
)

func main() {

	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("server start with fatal error , stop it , error =  , %s  , stack = %s", err, debug.Stack())
			os.Exit(1)
		}
	}()


	// 获取程序参数
	if len(os.Args) < 2 {
		fmt.Println("config file is empty")
		os.Exit(1)
	}
	configFile := os.Args[1]
	if "" == configFile {
		fmt.Println("config file is not found")
		os.Exit(1)
	}
	// 加载资源
	// 加载配置文件
	util.LoadConfiguration(configFile)
	util.LoadDbConfig(util.ServerConfig.DBConfig)

	// 配置web
	host := util.ServerConfig.Host
	port := util.ServerConfig.Port
	log.Printf("Server Start [ %v:%v ]", host, port)
	// 配置web路由
	http.HandleFunc("/payAck", handler.XyPayHandler)
	// 启动web
	log.Fatal(http.ListenAndServe(host + ":" + strconv.Itoa(port), nil))
}

//func getCurrentDirectory() string {
//	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
//	if err != nil {
//		log.Fatal(err)
//	}
//	return strings.Replace(dir, "\\", "/", -1)
//}