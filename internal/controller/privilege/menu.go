package privilege

import (
	"charon-janus/api/privilege/menu"
	"charon-janus/internal/service"
	"context"
)

var (
	Menu = sMenu{}
)

type sMenu struct{}

func (m *sMenu) List(ctx context.Context, req *menu.ListReq) (res *menu.ListRes, err error) {
	records, err := service.Menu().List(ctx, req.PlatFormCode)
	res = &menu.ListRes{
		Records: records,
	}
	return
}

func (m *sMenu) Edit(ctx context.Context, req *menu.EditReq) (_ *menu.EditRes, err error) {
	err = service.Menu().Edit(ctx, &req.MenuInput)
	return
}
