package menu

import (
	"charon-janus/internal/model/input"
	"github.com/gogf/gf/v2/frame/g"
)

type ListReq struct {
	g.Meta       `path:"/menu/list" method:"GET" summary:"获取菜单列表"`
	PlatFormCode string `json:"platForm_code"`
}

type ListRes struct {
	Records []input.MenuModelList `json:"records"`
}
