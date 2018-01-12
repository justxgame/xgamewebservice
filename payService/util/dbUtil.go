package util

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"runtime/debug"
)

func init() {

}

var dataSourceName string

func LoadDbConfig(dbConfig DBConfig) {
	// 如果加载配置文件失败，则退出系统
	defer func() {
		if err := recover(); err != nil {
			log.Fatalf("load db failed , exit system, err = %s , debug = %s ", err, debug.Stack())
		}
	}()

	// "root:root@tcp(localhost:3306)/?charset=utf8"
	dataSourceName = fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8&parseTime=True",
		dbConfig.UserName,
		dbConfig.Password,
		dbConfig.Address,
		dbConfig.DataBase)
	log.Printf("[db] dataSourceName = %v \n", dataSourceName)
	// test connection
	testConnection()
}

/**
 *  db 测试连接
 */
func testConnection() {
	defer func() {
		if err := recover(); err != nil {
			log.Panic("test connection failed ,", err)
		}
	}()

	db := GetDbConnection()
	defer db.Close()

	rows, err := db.Query("select 1 as testValue")
	defer rows.Close()

	if nil != err {
		log.Panic("db test connection error , %s  \n", err)
	}

	var testValue int

	if !rows.Next() {
		log.Panic("db test connection error , data is empty")
	}

	err = rows.Scan(&testValue)
	if err != nil {
		log.Panic("db test connection error %s \n", err)
	}

	if testValue != 1 {
		log.Panic("db test value  error , value = %v", testValue)
	}

}

/**
 *   只获取一行数据
 */
func GetOneRowData(rs *sql.Rows, dest ...interface{}) error {
	defer rs.Close()
	if rs.Next() {
		err := rs.Scan(dest...)
		return err

	} else {
		return nil
	}

}

//TODO thread pool
func GetDbConnection() *sql.DB {
	db, err := sql.Open("mysql", dataSourceName)
	if nil != err {
		log.Panic("db open error retun nil, datasource = %v , %s \n", dataSourceName, err)
	}
	return db
}
