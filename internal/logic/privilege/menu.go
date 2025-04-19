package privilege

import (
	"charon-janus/internal/dao"
	"charon-janus/internal/model/input"
	"charon-janus/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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
	if err = dao.AuthMenu.Ctx(ctx).Where(cols.PlatformCode, code).OrderDesc(cols.Order).Scan(&records); err != nil {
		g.Log().Error(ctx, err)
		return
	}
	return
}

func (s *sMenu) Edit(ctx context.Context, inp *input.MenuInput) (err error) {
	var (
		cols = dao.AuthMenu.Columns()
	)

	err = s.verify(ctx, inp.Id, inp.PlatformCode, g.Map{
		cols.Name:  inp.Name,
		cols.Path:  inp.Path,
		cols.Title: inp.Title,
	})

	if err != nil {
		return err
	}

	if inp.Id == 0 {
		_, err = dao.AuthMenu.Ctx(ctx).Data(&inp.AuthMenu).Insert()
		return
	}
	_, err = dao.AuthMenu.Ctx(ctx).WherePri(inp.Id).Data(&inp.AuthMenu).Update()
	return
}

func (s *sMenu) verify(ctx context.Context, id int, code string, scoreMap g.Map) (err error) {
	var (
		cols   = dao.AuthMenu.Columns()
		msgMap = g.MapStrStr{
			cols.Name:  "菜单名称已存在，请换一个",
			cols.Path:  "菜单路径已存在，请换一个",
			cols.Title: "菜单标题已存在，请换一个",
		}
	)

	for k, v := range scoreMap {
		count, err := dao.AuthMenu.Ctx(ctx).WhereNot(cols.Id, id).Where(cols.PlatformCode, code).Where(k, v).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return gerror.New(msgMap[k])
		}
	}
	return
}
