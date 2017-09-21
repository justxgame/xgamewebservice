package db

import "xgamewebservice/singleweb/util"

//import "xgamewebservice/singleweb/util"

/**
  server_info
 */

type ServerInfoModel struct {
	Server_id      string
	GameServerIp   string
	GameServerPort string
}

// server_id 查询 address
func QueryServerIpByServerId(server_id string) (ServerInfoModel, error) {
	var serverInfo ServerInfoModel
	db := util.GetDbConnection()
	defer db.Close()

	rows, err := db.Query("select server_id as Server_id , ip as GameServerIp ," +
		" gm_port as GameServerPort from server_info where server_id=?", server_id)
	if nil != err {
		return serverInfo, err
	}
	return serverInfo, util.GetOneRowData(rows,
		&serverInfo.Server_id,
		&serverInfo.GameServerIp,
		&serverInfo.GameServerPort)
}