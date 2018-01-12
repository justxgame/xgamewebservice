package main

import (
	"fmt"
	"log"
	"runtime/debug"
	"xgamewebservice/payService/db"
	"xgamewebservice/payService/util"
	//"strings"
)

const configFile = "/Users/william/work/go/src/xgamewebservice/singleweb/config.json"

func main() {

	// 加载资源
	// 加载配置文件
	util.LoadConfiguration(configFile)
	util.LoadDbConfig(util.ServerConfig.DBConfig)

	testPayBusiness()
}

func testPayBusiness() {

	// getPayLog
	//payLog , err := db.GetPayLog("bbb")
	//log.Printf("payLog= %s",payLog)
	//log.Printf("err= %s",err)

	// insert
	//var payLogDto db.PayLogDto
	//payLogDto.Order_id="ccc"
	//payLogDto.State=1
	//payLogDto.Pay_log = "{xxx=1}"
	//err := db.InsertPayLog(payLogDto)
	//
	//log.Print(strings.Contains(err.Error(),"Duplicate"))
	//log.Printf("err= %s",err)

	// update
	var payLogDto db.PayLogDto
	payLogDto.Order_id = "ccc"
	payLogDto.State = 3
	err := db.UpdatePayLogStateByOrderId(payLogDto)
	log.Printf("err= %s", err)

}

func testServerBusiness() {
	var serverInfo, err = db.QueryServerIpByServerId("1000")
	if nil != err {
		fmt.Printf("query failed , %s , %s ", err, debug.Stack())
	} else {
		fmt.Printf("serverInfo: %s ", serverInfo)
	}
}
