package input

import "charon-janus/internal/model/entity"

type LogInput struct {
	UserName string `json:"username" dc:"用户名"`
	Path     string `json:"path" dc:"请求路径"`
	ClientIp string `json:"client_ip" v:"ipv4" dc:"请求的client_ip"`
	PageReq
}

type LogRecords struct {
	entity.SysLog
}
