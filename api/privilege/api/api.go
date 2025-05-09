package api

import (
	"charon-janus/internal/model/input"
	"github.com/gogf/gf/v2/frame/g"
)

type ListReq struct {
	g.Meta       `path:"/privilege/api/list" method:"GET" dc:"接口列表"`
	PlatFormCode string `json:"platForm_code"`
}

type ListRes struct {
	Records []input.ApiModelList `json:"records"`
}

type EditReq struct {
	g.Meta `path:"/privilege/api/edit" method:"POST" dc:"编辑接口"`
	input.ApiInput
}
type EditRes struct{}
