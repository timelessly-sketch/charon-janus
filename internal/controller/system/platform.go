package system

import (
	"charon-janus/api/system/platform"
	"charon-janus/internal/service"
	"context"
)

var (
	PlatForm = cPlatForm{}
)

type cPlatForm struct{}

func (p *cPlatForm) List(ctx context.Context, req *platform.ListReq) (res *platform.ListRes, err error) {
	records, total, err := service.PlatForm().List(ctx, &req.PlatFormInput)
	res = &platform.ListRes{
		Records: records,
		Total:   total,
	}
	return
}
