package privilege

import (
	"charon-janus/api/privilege/api"
	"charon-janus/internal/service"
	"context"
)

var (
	Api = cApi{}
)

type cApi struct{}

func (a *cApi) List(ctx context.Context, req *api.ListReq) (res *api.ListRes, err error) {
	records, err := service.Api().List(ctx, req.PlatFormCode)
	res = &api.ListRes{
		Records: records,
	}
	return
}

func (a *cApi) Edit(ctx context.Context, req *api.EditReq) (_ *api.EditRes, err error) {
	err = service.Api().Edit(ctx, req.ApiInput)
	return
}
