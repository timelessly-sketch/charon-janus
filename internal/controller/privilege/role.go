package privilege

import (
	"charon-janus/api/privilege/role"
	"charon-janus/internal/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Role = sRole{}
)

type sRole struct{}

func (r *sRole) List(ctx context.Context, req *role.ListReq) (res *role.ListRes, err error) {
	records, err := service.Role().List(ctx, req.PlatFormCode)
	res = &role.ListRes{
		Records: records,
	}
	return
}

func (r *sRole) Edit(ctx context.Context, req *role.EditReq) (res *role.EditRes, err error) {
	if err = service.Role().Edit(ctx, &req.RoleEditInput); err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

func (r *sRole) Detail(ctx context.Context, req *role.DetailReq) (res *role.DetailRes, err error) {
	records, err := service.Role().Detail(ctx, req.Id)
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	res = &role.DetailRes{
		RoleDetailList: records,
	}
	return
}
