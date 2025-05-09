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

type sApi struct{}

func NewApi() *sApi {
	return &sApi{}
}

func init() {
	service.RegisterApi(NewApi())
}

func (s *sApi) List(ctx context.Context, code string) (records []input.ApiModelList, err error) {
	err = dao.AuthApi.Ctx(ctx).Where(dao.AuthApi.Columns().PlatformCode, code).Scan(&records)
	return
}

func (s *sApi) Edit(ctx context.Context, inp input.ApiInput) (err error) {
	cols := dao.AuthApi.Columns()

	err = s.verify(ctx, inp.Id, inp.PlatformCode, g.Map{
		cols.Name:  inp.Name,
		cols.Path:  inp.Path,
		cols.Title: inp.Title,
	})
	if err != nil {
		return err
	}

	if inp.Id == 0 {
		_, err = dao.AuthApi.Ctx(ctx).Data(&inp).Insert()
		return
	}
	err = g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = dao.AuthApi.Ctx(ctx).WherePri(inp.Id).Data(&inp).Update(); err != nil {
			return err
		}
		return cache.RemoveByPrefix(ctx, consts.LoginApi)
	})
	return
}

func (s *sApi) verify(ctx context.Context, id int, code string, scoreMap g.Map) (err error) {
	var (
		cols   = dao.AuthApi.Columns()
		msgMap = g.MapStrStr{
			cols.Name:  "接口名称已存在，请换一个",
			cols.Path:  "接口路径已存在，请换一个",
			cols.Title: "接口标题已存在，请换一个",
		}
	)
	g.Log().Info(ctx, scoreMap)
	for k, v := range scoreMap {
		if k == cols.Path && v == "" {
			continue
		}
		count, err := dao.AuthApi.Ctx(ctx).WhereNot(cols.Id, id).Where(cols.PlatformCode, code).Where(k, v).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return gerror.New(msgMap[k])
		}
	}
	return
}

func (s *sApi) DetailApi(ctx context.Context, code string, id int) (records []input.ApiModelList, ids []int) {
	var (
		cols    = dao.AuthApi.Columns()
		colsApi = dao.AuthRoleApi.Columns()
	)

	_ = dao.AuthApi.Ctx(ctx).Where(cols.PlatformCode, code).Scan(&records)
	array, _ := dao.AuthRoleApi.Ctx(ctx).Fields(colsApi.ApiId).Where(colsApi.RoleId, id).Array()

	return records, gvar.New(array).Ints()
}

func (s *sApi) RoleApiEdit(ctx context.Context, apiIds []int, roleId int) (err error) {
	var (
		colsApi = dao.AuthRoleApi.Columns()
		apiList = make([]entity.AuthRoleApi, 0)
	)

	array, _ := dao.AuthRoleApi.Ctx(ctx).Fields(colsApi.ApiId).Where(colsApi.RoleId, roleId).Array()
	addedMenu, removedMenu := convert.Contrast(gvar.New(array).Ints(), apiIds)
	for _, id := range addedMenu {
		apiList = append(apiList, entity.AuthRoleApi{
			RoleId: roleId,
			ApiId:  id,
		})
	}
	if len(removedMenu) > 0 {
		if _, err = dao.AuthRoleApi.Ctx(ctx).Where(colsApi.RoleId, roleId).WhereIn(colsApi.ApiId, removedMenu).Delete(); err != nil {
			return
		}
	}
	if len(addedMenu) > 0 {
		_, err = dao.AuthRoleApi.Ctx(ctx).Data(&apiList).Insert()
	}
	return
}
