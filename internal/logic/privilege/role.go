package privilege

import (
	"charon-janus/internal/dao"
	"charon-janus/internal/library/cache"
	"charon-janus/internal/model/entity"
	"charon-janus/internal/model/input"
	"charon-janus/internal/service"
	"charon-janus/utility/convert"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type sRole struct{}

func NewRole() *sRole {
	return &sRole{}
}

func init() {
	service.RegisterRole(NewRole())
}

func (s *sRole) List(ctx context.Context, code string) (records []input.RoleModelList, err error) {
	cols := dao.AuthRole.Columns()
	if err = dao.AuthRole.Ctx(ctx).Where(cols.PlatformCode, code).OrderAsc(cols.RoleSort).Scan(&records); err != nil {
		return
	}
	return
}

func (s *sRole) Edit(ctx context.Context, inp *input.RoleEditInput) (err error) {
	var (
		cols     = dao.AuthRole.Columns()
		authRols = dao.AuthRoleMenu.Columns()
		menuList = make([]entity.AuthRoleMenu, 0)
	)

	err = s.verify(ctx, inp.Id, g.Map{
		cols.RoleKey: inp.RoleKey,
	})

	if err != nil {
		return
	}
	if inp.Id == 0 {
		return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			return
		})
	}
	array, err := cache.Instance().Get(ctx, s.cacheKey(inp.Id))
	if err != nil {
		return
	}
	added, removed := convert.Contrast(gvar.New(array).Ints(), inp.MenuIds)
	for _, id := range added {
		menuList = append(menuList, entity.AuthRoleMenu{
			RoleId: inp.Id,
			MenuId: id,
		})
	}
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if len(removed) > 0 {
			if _, err = dao.AuthRoleMenu.Ctx(ctx).Where(authRols.RoleId, inp.Id).WhereIn(authRols.MenuId, removed).Delete(); err != nil {
				return err
			}
		}
		if len(added) > 0 {
			_, err = dao.AuthRoleMenu.Ctx(ctx).Data(&menuList).Insert()
		}
		return
	})
}

func (s *sRole) Detail(ctx context.Context, id int) (records input.RoleDetailList, err error) {
	var (
		role     = entity.AuthRole{}
		cols     = dao.AuthRoleMenu.Columns()
		menuList []entity.AuthMenu
	)
	if err = dao.AuthRole.Ctx(ctx).WherePri(id).Scan(&role); err != nil {
		return
	}

	if err = dao.AuthMenu.Ctx(ctx).Where(dao.AuthMenu.Columns().PlatformCode, role.PlatformCode).Scan(&menuList); err != nil {
		return
	}
	array, _ := dao.AuthRoleMenu.Ctx(ctx).Fields(cols.MenuId).Where(cols.RoleId, role.Id).Array()
	records = input.RoleDetailList{
		MenuList: menuList,
		MenuIds:  gvar.New(array).Ints(),
	}
	_ = cache.Instance().Set(ctx, s.cacheKey(id), array, 24*time.Hour)
	return
}

func (s *sRole) verify(ctx context.Context, id int, scoreMap g.Map) (err error) {
	var (
		cols   = dao.AuthRole.Columns()
		msgMap = g.MapStrStr{
			cols.RoleName: "角色名称已存在，请换一个",
			cols.RoleKey:  "角色编码已存在，请换一个",
		}
	)

	for k, v := range scoreMap {
		count, err := dao.AuthRole.Ctx(ctx).WhereNot(cols.Id, id).Where(k, v).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return gerror.New(msgMap[k])
		}
	}
	return
}

func (s *sRole) cacheKey(id int) string {
	return fmt.Sprintf("%s:%d", "RoleDetail", id)
}
