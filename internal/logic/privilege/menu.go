package privilege

import (
	"charon-janus/internal/consts"
	"charon-janus/internal/dao"
	"charon-janus/internal/library/cache"
	"charon-janus/internal/model/entity"
	"charon-janus/internal/model/input"
	"charon-janus/internal/service"
	"charon-janus/utility/convert"
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
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
	if err = dao.AuthMenu.Ctx(ctx).Where(cols.PlatformCode, code).OrderAsc(cols.Order).Scan(&records); err != nil {
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
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = dao.AuthMenu.Ctx(ctx).WherePri(inp.Id).Data(&inp.AuthMenu).Update(); err != nil {
			return err
		}
		return cache.RemoveByPrefix(ctx, consts.LoginMenu)
	})
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

func (s *sMenu) DetailMenu(ctx context.Context, code string, id int) (records []input.MenuModelList, ids []int) {
	var (
		cols     = dao.AuthMenu.Columns()
		colsMenu = dao.AuthRoleMenu.Columns()
	)

	_ = dao.AuthMenu.Ctx(ctx).Where(cols.PlatformCode, code).OrderAsc(cols.Order).Scan(&records)
	array, _ := dao.AuthRoleMenu.Ctx(ctx).Fields(colsMenu.MenuId).Where(colsMenu.RoleId, id).Array()

	return records, gvar.New(array).Ints()
}

func (s *sMenu) RoleMenuEdit(ctx context.Context, menuIds []int, roleId int) (err error) {
	var (
		colsMenu = dao.AuthRoleMenu.Columns()
		menuList = make([]entity.AuthRoleMenu, 0)
	)

	array, _ := dao.AuthRoleMenu.Ctx(ctx).Fields(colsMenu.MenuId).Where(colsMenu.RoleId, roleId).Array()
	addedMenu, removedMenu := convert.Contrast(gvar.New(array).Ints(), menuIds)
	for _, id := range addedMenu {
		menuList = append(menuList, entity.AuthRoleMenu{
			RoleId: roleId,
			MenuId: id,
		})
	}
	if len(removedMenu) > 0 {
		if _, err = dao.AuthRoleMenu.Ctx(ctx).Where(colsMenu.RoleId, roleId).WhereIn(colsMenu.MenuId, removedMenu).Delete(); err != nil {
			return
		}
	}
	if len(addedMenu) > 0 {
		_, err = dao.AuthRoleMenu.Ctx(ctx).Data(&menuList).Insert()
	}
	return
}
