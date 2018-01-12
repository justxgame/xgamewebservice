package db

import (
	"log"
	"time"
	"xgamewebservice/payService/util"
)

//import "xgamewebservice/singleweb/util"

/**
  server_info
*/
const (
	PAY_INIT    = 1
	PAY_FAILED  = 2
	PAY_SUCCESS = 9
)

type PayLogDto struct {
	Order_id         string
	Pay_log          string
	State            int
	Create_Time      time.Time
	Last_update_time time.Time
	Msg              string
}

// server_id 查询 address TODO 暂时不用
func GetPayLog(order_id string) (PayLogDto, error) {
	var payLog PayLogDto
	db := util.GetDbConnection()
	defer db.Close()
	rows, err := db.Query("select order_id as Order_id , pay_log as Pay_log, "+
		" state as State , create_time as Create_Time , last_update_time as Last_update_time "+
		" from pay_log where order_id=? ", order_id)
	if nil != err {
		return payLog, err
	}
	return payLog, util.GetOneRowData(rows,
		&payLog.Order_id,
		&payLog.Pay_log,
		&payLog.State,
		&payLog.Create_Time,
		&payLog.Last_update_time)
}

// 如果order存在，报错
func InsertPayLog(payLogDto PayLogDto) error {
	db := util.GetDbConnection()
	defer db.Close()
	res, err := db.Exec("insert into pay_log (order_id,pay_log,state,create_time,last_update_time) "+
		"VALUES (?,?,?,NOW(),NOW())", payLogDto.Order_id, payLogDto.Pay_log, payLogDto.State)

	log.Printf("insert result : %s ", res)
	if nil != err {
		return err
	} else {
		return nil
	}
}

/**
只跟新状态和msg
*/
func UpdatePayLogStateByOrderId(payLogDto PayLogDto) error {
	db := util.GetDbConnection()
	defer db.Close()
	res, err := db.Exec("update pay_log set state=? , msg =?, last_update_time=NOW() where order_id =?", payLogDto.State, payLogDto.Msg, payLogDto.Order_id)
	log.Printf("upate result : %s ", res)
	if nil != err {
		return err
	} else {
		return nil
	}
}

// 删除 TODO
func DeletePayLogByOrderId(payLogDto PayLogDto) error {
	db := util.GetDbConnection()
	defer db.Close()
	res, err := db.Exec("delet from pay_log where order_id =?", payLogDto.Order_id)
	log.Printf("delete result : %s ", res)
	if nil != err {
		return err
	} else {
		return nil
	}
}
