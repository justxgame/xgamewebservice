package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/?charset=utf8")

	defer db.Close()

	if nil != err {
		fmt.Printf("db open error %s \n", err)
	}

	rows, err := db.Query("select user_id from adv_dsp.user_config")

	defer rows.Close()
	if nil != err {
		fmt.Printf("query err %s \n", err)
	}

	type User_Config struct {
		user_id          string
		user_name        string
		user_password    string
		approve          int
		is_administrator int
	}

	var user_config User_Config

	for rows.Next() {
		err := rows.Scan(&user_config.user_id)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("user_id %v \n", user_config)
	}
}

func fck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/**
  using a map
*/
type mapStringScan struct {
	// cp are the column pointers
	cp []interface{}
	// row contains the final result
	row      map[string]string
	colCount int
	colNames []string
}

func NewMapStringScan(columnNames []string) *mapStringScan {
	lenCN := len(columnNames)
	s := &mapStringScan{
		cp:       make([]interface{}, lenCN),
		row:      make(map[string]string, lenCN),
		colCount: lenCN,
		colNames: columnNames,
	}
	for i := 0; i < lenCN; i++ {
		s.cp[i] = new(sql.RawBytes)
	}
	return s
}
