package input

import "charon-janus/internal/model/entity"

type LogInput struct {
	UserName string `json:"username" dc:"用户名"`
	Method   string `json:"method" dc:"请求方法"`
	ClientIp string `json:"client_ip" dc:"请求的client_ip"`
	StartAt  string `json:"start_at"`
	EndAt    string `json:"end_at"`
	PageReq
}

type LogRecords struct {
	Records []entity.SysLog
}
