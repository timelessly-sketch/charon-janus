package platform

import (
	"charon-janus/internal/model/input"
	"github.com/gogf/gf/v2/frame/g"
)

type ListReq struct {
	g.Meta `path:"/platformList" method:"get" summary:"获取平台列表"`
	input.PlatFormInput
}

type ListRes struct {
	Records []input.PlatFormModelList `json:"records"`
	Total   int                       `json:"total"`
}
