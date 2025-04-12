package privilege

import (
	"charon-janus/internal/dao"
	"charon-janus/internal/model/input"
	"charon-janus/internal/service"
	"context"
)

type sMenu struct{}

func NewMenu() *sMenu {
	return &sMenu{}
}

func init() {
	service.RegisterMenu(NewMenu())
}

func (s *sMenu) List(ctx context.Context, code string) (records []input.MenuModelList, err error) {
	cols := dao.AuthMenu.Columns()
	err = dao.AuthMenu.Ctx(ctx).Where(cols.PlatformCode, code).OrderDesc(cols.Order).Scan(&records)
	return
}
