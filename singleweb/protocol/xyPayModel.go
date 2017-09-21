package protocol

const (
	XY_SUCCESS = 0  //发货成功
	XY_PARAMETER_ERROR = 1  //参数错误
	XY_USER_NOT_EXIST_ERROR = 2 //玩家不存在
	XY_SERVER_NOT_EXIST_ERROR = 3 //游戏服不存在
	XY_ORDER_IS_EXIST_ERROR = 4 //订单已存在
	XY_INFO_ERROR = 5 //透传信息错误
	XY_SIGN_ERROR = 6 //签名校验错误
	XY_DB_ERROR = 7 //数据库错误
	XY_OTHERS_ERROR = 8 //其它错误
)


// 支付callback确认信息
type XyPayAckModel struct {
	Orderid  string         `json:"orderid"`  // xyzs 平台订单号
	Uid      string         `json:"uid"`      // xyzs 平台用户 ID
	Serverid string         `json:"serverid"` //serverid
	Amount   string        `json:"amount"`    //点型 人民币消耗金额，单位:元
	Extra    string         `json:"extra"`
	Ts       string         `json:"ts"`
	Sign     string         `json:"sign"`
	Sig      string         `json:"sig"`
						  // 游戏方接受的字段
	Pid      string                        `json:"pid"`
	PayType  string                        `json:"pay_type"`
}

// 游戏方接受的字段
type XyPayServerAckModel struct {
	Orderid  string   `json:"order_id"`    // xyzs 平台订单号
	Serverid uint32   `json:"server_id"`   //serverid
	Pid      uint64         `json:"pid"`
	PayType  int32         `json:"pay_type"`
	Money    string         `json:"money"` //游戏方暂时不处理
}

// xy的callback返回值
type XyPayResponseModel  struct {
	Code    int `json:"code"`
	Message string `json:"message"`
}

// 游戏的返回值
type XyPayResponseGameServerModel  struct {
	Code int        `json:"Code"`
	Desc string  `json:"Desc"`
}


