package system

import (
	"charon-janus/api/system/log"
	"charon-janus/internal/service"
	"context"
)

var (
	Log = sLog{}
)

type sLog struct{}

func (s sLog) List(ctx context.Context, req *log.ListReq) (res *log.ListRes, err error) {
	records, total, err := service.Log().List(ctx, &req.LogInput)
	if err != nil {
		return
	}
	res = &log.ListRes{
		Total:   total,
		Records: records,
	}
	return
}
