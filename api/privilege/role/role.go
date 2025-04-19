package role

import (
	"charon-janus/internal/model/input"
	"github.com/gogf/gf/v2/frame/g"
)

type ListReq struct {
	g.Meta       `path:"/role/list" method:"GET" summary:"角色列表"`
	PlatFormCode string `json:"platFormCode"`
}

type ListRes struct {
	Records []input.RoleModelList `json:"records"`
}

type EditReq struct {
	g.Meta `path:"/role/edit" method:"POST" summary:"编辑角色"`
	input.RoleEditInput
}

type EditRes struct{}

type DetailReq struct {
	g.Meta `path:"/role/detail/{id}" method:"get" summary:"获取当前规则详情"`
	Id     int `json:"id" v:"required#规则ID缺失"`
}

type DetailRes struct {
	input.RoleDetailList
}
