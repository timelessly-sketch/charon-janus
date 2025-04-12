package platform

import (
	"charon-janus/internal/model/input"
	"github.com/gogf/gf/v2/frame/g"
)

type ListReq struct {
	g.Meta `path:"/platform/list" method:"get" summary:"获取平台列表"`
	input.PageReq
}

type ListRes struct {
	Records []input.PlatFormModelList `json:"records"`
	Total   int                       `json:"total"`
}

type EditReq struct {
	g.Meta `path:"platform/edit" method:"post" summary:"编辑平台"`
	input.PlatFormEditInput
}

type EditRes struct{}

type OptionsReq struct {
	g.Meta `path:"platform/options" method:"get" summary:"获取用户平台权限"`
}

type OptionsRes struct {
	Records []input.PlatFormModelList `json:"records"`
}
