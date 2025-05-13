package login

import (
	"charon-janus/api/login"
	"charon-janus/internal/service"
	"context"
)

var (
	Login = sLogin{}
)

type sLogin struct{}

func (s sLogin) Login(ctx context.Context, req *login.AccountLoginReq) (res *login.AccountLoginRes, err error) {
	records, err := service.Login().Login(ctx, &req.AccountLoginInp)
	defer service.Login().InsertLoginLog(ctx, records, err)
	if err != nil {
		return
	}
	res = &login.AccountLoginRes{
		LoginModel: records,
	}
	return
}

func (s sLogin) Routers(ctx context.Context, req *login.RoutesReq) (res *login.RoutesRes, err error) {
	records, err := service.Login().UserRoutes(ctx, req.PlatFormCode)
	res = &login.RoutesRes{
		Records: records,
	}
	return
}

func (s sLogin) LoginLogList(ctx context.Context, req *login.ListReq) (res *login.ListRes, err error) {
	records, total, err := service.Login().GetLoginLog(ctx, req.LoginLogInp)
	res = &login.ListRes{
		Records: records,
		Total:   total,
	}
	return
}
