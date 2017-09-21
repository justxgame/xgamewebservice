package handler

import (
	"io/ioutil"
	"net/http"
	"log"
	"encoding/json"
	"xgamewebservice/singleweb/protocol"
	"xgamewebservice/singleweb/signature"
	"xgamewebservice/singleweb/util"
	"strings"
	"xgamewebservice/singleweb/db"
	"strconv"
	"runtime/debug"
)

/**
orderid = 100031065_45805054_1505997333_1445
uid = 45805054
serverid = 1000
amount = 0.01
extra = 6|1
ts = 1505997333
sign = 6516e118065e34f5e6fff6f0fd34966d
sig = 55d5a44cbf02c403ea3a9b130f9f20fb
 */
func XyPayHandler(w http.ResponseWriter, r *http.Request) {

	// 随机生成支付顺序
	var requestSeq = util.RandStringBytesMaskImpr(16)

	// 最终异常捕获
	defer func() {
		if err := recover(); err != nil {
			log.Printf("[XyPayHandler-%s] failed=%s ,stack=%s ", requestSeq, err, debug.Stack())
			xyResponseWrap(w, protocol.XY_OTHERS_ERROR, "其他")
		}
		log.Printf("[XyPayHandler-%v] pay request end", requestSeq)
	}()

	defer r.Body.Close()
	log.Printf("[XyPayHandler-%v] pay request start", requestSeq)
	log.Printf("[XyPayHandler-%v] remote address=%s  ", requestSeq, r.RemoteAddr)
	// must post
	if r.Method != "POST" {
		log.Println("[XyPayHandler-%v] not post type , type=%v", requestSeq, r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var payAckModel protocol.XyPayAckModel
	err := r.ParseForm()
	if nil != err {
		log.Printf("[XyPayHandler-%v] parse form error=%s ", requestSeq, err)
	}
	log.Printf("[XyPayHandler-%v] request parameters=%s", requestSeq, r.PostForm)

	// 表单解析
	payAckModel.Orderid = r.PostFormValue("orderid")
	payAckModel.Uid = r.PostFormValue("uid")
	payAckModel.Amount = r.PostFormValue("amount")
	payAckModel.Serverid = r.PostFormValue("serverid")
	payAckModel.Extra = r.PostFormValue("extra")
	payAckModel.Ts = r.PostFormValue("ts")
	payAckModel.Sign = r.PostFormValue("sign")
	payAckModel.Sig = r.PostFormValue("sig")
	//组装游戏方需要的字段
	var splitArray []string
	if (payAckModel.Extra != "") {
		splitArray = strings.Split(payAckModel.Extra, "|")
	}
	if (splitArray != nil && len(splitArray) >= 2 &&isNumber(splitArray[0])&&isNumber(splitArray[1])) {
		payAckModel.Pid = splitArray[0]
		payAckModel.PayType = splitArray[1]
	} else {
		log.Printf("[XyPayHandler-%v] extra parse error , value=$v", requestSeq, payAckModel.Extra)
		var response protocol.XyPayResponseModel
		response.Code = 1
		returnBody, _ := json.Marshal(&response)
		w.Write(returnBody)
		return
	}

	// 组装成游戏方需要的对象
	var payServerAckModel protocol.XyPayServerAckModel
	payServerAckModel.Orderid = payAckModel.Orderid
	a, _ := strconv.ParseInt(payAckModel.PayType, 10, 32)
	payServerAckModel.PayType = int32(a)
	b, _ := strconv.ParseUint(payAckModel.Pid, 10, 32)
	payServerAckModel.Pid = b
	c, _ := strconv.ParseUint(payAckModel.Serverid, 10, 32)
	payServerAckModel.Serverid = uint32(c)
	payServerAckModel.Money = payAckModel.Amount

	// 验证
	if (!signPass(payAckModel, requestSeq)) {
		log.Printf("[XyPayHandler-%v] sign failed !", requestSeq)
		xyResponseWrap(w, protocol.XY_SIGN_ERROR, "验证错误")
		return
	} else {
		log.Printf("[XyPayHandler-%v] sign success !", requestSeq)
	}


	// 获取转发地址
	if (payAckModel.Serverid == "") {
		log.Panic("serverid is empty")
	}

	serverModel, err := db.QueryServerIpByServerId(payAckModel.Serverid)
	if nil != err {
		log.Panic("query db error")
	}

	// 转发host
	transferHost := "http://" + serverModel.GameServerIp + ":" + serverModel.GameServerPort + "/pay"
	log.Printf("[XyPayHandler-%v] transfer host=%v ", requestSeq, transferHost)

	payServerData, _ := json.Marshal(payServerAckModel)
	log.Printf("[XyPayHandler-%v] transfer data=%v ", requestSeq, string(payServerData))
	resp, err := http.Post(transferHost, "application/json", strings.NewReader(string(payServerData)))

	if (nil != err) {
		log.Printf("[XyPayHandler-%v]  transfer failed ,s %s ", requestSeq, err)
		xyResponseWrap(w, protocol.XY_OTHERS_ERROR, "转发错误")
		return
	}

	if (nil == resp) {
		log.Printf("[XyPayHandler-%v] transfer resp is nil ", requestSeq)
		xyResponseWrap(w, protocol.XY_OTHERS_ERROR, "转发错误")
		return
	}

	resp_body, _ := ioutil.ReadAll(resp.Body)
	if http.StatusOK == resp.StatusCode {
		// todo parse resp
		log.Printf("[XyPayHandler-%v] receive game server response body=%v ", requestSeq, string(resp_body))
		var responseGameServerModel protocol.XyPayResponseGameServerModel
		err = json.Unmarshal(resp_body, &responseGameServerModel)
		if (err != nil ) {
			log.Printf("[XyPayHandler-%v] receive game server parse failed , resp_body=%v ", requestSeq, string(resp_body))
		}
		if (responseGameServerModel.Code == 0) {
			// 成功
			var response protocol.XyPayResponseModel
			response.Code = 0
			returnBody, _ := json.Marshal(&response)
			w.Write(returnBody)
			log.Printf("[XyPayHandler-%v] success ", requestSeq)
			return
		} else {
			log.Printf("[XyPayHandler-%v] failed code=%v ", requestSeq, responseGameServerModel.Code)
		}
	} else {
		log.Printf("[XyPayHandler-%v] failed code=%s , body=%s", requestSeq, resp.StatusCode, string(resp_body))
	}

	// final failed
	xyResponseWrap(w, protocol.XY_OTHERS_ERROR, "其他错误")
	return

}

func signPass(payAckModel protocol.XyPayAckModel, requestSeq string) bool {
	// 验证
	var m = make(map[string]string)
	m["orderid"] = payAckModel.Orderid
	m["uid"] = payAckModel.Uid
	m["amount"] = payAckModel.Amount
	m["serverid"] = payAckModel.Serverid
	m["extra"] = payAckModel.Extra
	m["ts"] = payAckModel.Ts
	m["sign"] = payAckModel.Sign
	m["sig"] = payAckModel.Sig
	//验证App签名串
	appSign := signature.GetGenSafeSign(m, util.ServerConfig.AppKey)
	log.Printf("[XyPayHandler-%v] appsign=%s", requestSeq, appSign)
	// 验证错误
	if m["sign"] == "" || appSign != m["sign"] {
		return false
	}
	//如果支付签名串存在就验证
	if (m["sig"] != "") {
		var paySign = signature.GetGenSafeSign(m, util.ServerConfig.PayKey)
		log.Printf("[XyPayHandler-%v] paysign=%s", requestSeq, appSign)
		// 验证错误
		if paySign != m["sig"] {
			return false
		}
	}
	return true
}

/**
包装返回值
 */
func xyResponseWrap(w http.ResponseWriter, code int, message string) {
	var response protocol.XyPayResponseModel
	response.Code = code
	response.Message = message
	returnBody, _ := json.Marshal(&response)
	w.Write(returnBody)
}

func isNumber(v string) bool {
	if _, err := strconv.Atoi(v); err == nil {
		return true;
	} else {
		return false;
	}
}


