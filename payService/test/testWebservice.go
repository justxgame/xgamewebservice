package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"xgamewebservice/payService/protocol"
	_ "xgamewebservice/payService/util"
)

var port = 8888

func main() {
	log.Printf("test Server Start !")
	//log.Info("info")
	http.HandleFunc("/", contentHandler) //	设置访问路由
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(port), nil))
}

func contentHandler(rw http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	log.Printf("====================================")
	log.Printf("Req: %s %s", r.Host, r.URL.Path)
	log.Printf("Type: %s ", r.Method)
	// read http body
	body, ioErr := ioutil.ReadAll(r.Body)
	log.Printf("body: %v", string(body))

	if ioErr != nil {
		log.Println("http io error", ioErr, string(body))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	var res protocol.XyPayResponseModel
	res.Code = 1
	res.Message = "ok"

	d, _ := json.Marshal(&res)

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(200)
	rw.Write(d)
	//json.NewEncoder(rw).Encode(&u)
}
