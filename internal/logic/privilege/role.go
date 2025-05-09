package privilege

import (
	"charon-janus/internal/dao"
	"charon-janus/internal/model/entity"
	"charon-janus/internal/model/input"
	"charon-janus/internal/service"
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
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
	cols := dao.AuthRole.Columns()
	if err = s.verify(ctx, inp.Id, g.Map{dao.AuthRole.Columns().RoleKey: inp.RoleKey}); err != nil {
		return
	}
	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if inp.Id == 0 {
			id, err := dao.AuthRole.Ctx(ctx).Data(g.Map{
				cols.RoleKey:      inp.RoleKey,
				cols.PlatformCode: inp.PlatformCode,
				cols.RoleName:     inp.RoleName,
				cols.RoleSort:     inp.RoleSort,
				cols.Status:       inp.Status,
			}).InsertAndGetId()
			if err != nil {
				return err
			}
			inp.Id = gconv.Int(id)
		} else {
			if _, err = dao.AuthRole.Ctx(ctx).WherePri(inp.Id).Data(g.Map{
				cols.RoleKey:      inp.RoleKey,
				cols.PlatformCode: inp.PlatformCode,
				cols.RoleName:     inp.RoleName,
				cols.RoleSort:     inp.RoleSort,
				cols.Status:       inp.Status,
			}).Update(); err != nil {
				return
			}
		}

		if err = service.Menu().RoleMenuEdit(ctx, inp.MenuIds, inp.Id); err != nil {
			return
		}

		if err = service.Api().RoleApiEdit(ctx, inp.ApiIds, inp.Id); err != nil {
			return
		}
		return
	})
}

func (s *sRole) Detail(ctx context.Context, id int) (records input.RoleDetailList, err error) {
	var (
		role = entity.AuthRole{}
	)
	if err = dao.AuthRole.Ctx(ctx).WherePri(id).Scan(&role); err != nil {
		return
	}
	menuList, arrayMenu := service.Menu().DetailMenu(ctx, role.PlatformCode, id)
	apiList, arrayApi := service.Api().DetailApi(ctx, role.PlatformCode, id)

	records = input.RoleDetailList{
		MenuList: menuList,
		MenuIds:  gvar.New(arrayMenu).Ints(),
		ApiList:  apiList,
		ApiIds:   gvar.New(arrayApi).Ints(),
	}
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
