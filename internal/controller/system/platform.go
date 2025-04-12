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
	records, total, err := service.PlatForm().List(ctx, &req.PageReq)
	res = &platform.ListRes{
		Records: records,
		Total:   total,
	}
	return
}

func (p *cPlatForm) Edit(ctx context.Context, req *platform.EditReq) (_ *platform.EditRes, err error) {
	err = service.PlatForm().Edit(ctx, &req.PlatFormEditInput)
	return
}

func (p *cPlatForm) Options(ctx context.Context, _ *platform.OptionsReq) (res *platform.OptionsRes, err error) {
	records, err := service.PlatForm().Options(ctx)
	res = &platform.OptionsRes{
		Records: records,
	}
	return
}
