package main

import (
	"fmt"
	"xgamewebservice/singleweb/signature"
)

var m = make(map[string]string)
var AppKey = "b6da5900312eaaae2b3f78b9073106f0"
var PayKey = "tdTmL9KWaNBwsW40FH7FVQKEYkx9UYfk";

func main() {

	var orderid = "100001_10001_1407811507_9347";
	var uid = "10001";
	var amount = "1.00";
	var serverid = "0";
	var extra = "201408036987";
	var ts = "1407811565";
	var sign = "6631e0616ac73df7ea610bb4f9e03e0d";
	var sig = "2957d2b348dec094c5d7f9fb84821d1c";

	m["orderid"] = orderid
	m["uid"] = uid
	m["amount"] = amount
	m["serverid"] = serverid
	m["extra"] = extra
	m["ts"] = ts
	m["sign"] = sign
	m["sig"] = sig

	//验证App签名串
	var appSign = signature.GetGenSafeSign(m, AppKey)
	fmt.Println("app签名", appSign)

	//如果支付签名串存在就验证
	if (m["sign"] != "") {
		var paySign = signature.GetGenSafeSign(m, PayKey)
		fmt.Println("pay签名",paySign)
	}

}
