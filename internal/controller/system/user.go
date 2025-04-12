package system

import (
	"charon-janus/api/system/user"
	"charon-janus/internal/service"
	"context"
)

var (
	User = sUser{}
)

type sUser struct{}

func (s *sUser) List(ctx context.Context, req *user.ListReq) (res *user.ListRes, err error) {
	records, total, err := service.User().List(ctx, &req.UserInput)
	res = &user.ListRes{
		Records: records,
		Total:   total,
	}
	return
}

func (s *sUser) Detail(ctx context.Context, req *user.DetailReq) (res *user.DetailRes, err error) {
	records, err := service.User().Detail(ctx, req.Id)
	res = &user.DetailRes{
		UserModelDetail: records,
	}
	return
}

func (s *sUser) Edit(ctx context.Context, req *user.EditReq) (res *user.EditRes, err error) {
	err = service.User().Edit(ctx, &req.UserEditInput)
	return
}

func (s *sUser) Reset(ctx context.Context, req *user.ResetReq) (_ *user.ResetRes, err error) {
	err = service.User().Reset(ctx, req.Username)
	return
}
