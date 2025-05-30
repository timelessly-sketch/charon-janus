package log

import (
	"charon-janus/internal/model/input"
	"github.com/gogf/gf/v2/frame/g"
)

type ListReq struct {
	g.Meta `path:"/system/log/list" method:"get"`
	input.LogInput
}

type ListRes struct {
	Total   int                `json:"total"`
	Records []input.LogRecords `json:"records"`
}
