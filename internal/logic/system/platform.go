package system

import (
	"charon-janus/internal/dao"
	"charon-janus/internal/library/cache"
	"charon-janus/internal/model/entity"
	"charon-janus/internal/model/input"
	"charon-janus/internal/service"
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sPlatForm struct{}

func NewPlatForm() *sPlatForm {
	return &sPlatForm{}
}

func init() {
	service.RegisterPlatForm(NewPlatForm())
}

func (s *sPlatForm) List(ctx context.Context, inp *input.PageReq) (records []input.PlatFormModelList, total int, err error) {
	err = dao.SysPlatform.Ctx(ctx).OrderDesc(dao.SysPlatform.Columns().PlatformSort).
		Page(inp.Page, inp.Size).ScanAndCount(&records, &total, true)
	return
}

func (s *sPlatForm) Edit(ctx context.Context, inp *input.PlatFormEditInput) (err error) {
	cols := dao.SysPlatform.Columns()
	if err = s.verify(ctx, inp.Id, g.Map{
		cols.PlatformCode: inp.PlatformCode,
		cols.ServerUrl:    inp.ServerUrl,
	}); err != nil {
		return err
	}
	if inp.Id == 0 {
		_, err = dao.SysPlatform.Ctx(ctx).OmitEmpty().Data(&inp.SysPlatform).Insert()
		return err
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = dao.SysPlatform.Ctx(ctx).WherePri(inp.Id).Data(&inp.SysPlatform).Update(); err != nil {
			return err
		}
		return cache.ClearDBCache(ctx)
	})
}

func (s *sPlatForm) Options(ctx context.Context) (records []input.PlatFormModelList, err error) {
	var (
		cols     = dao.SysAuthRoles.Columns()
		identity = service.Middleware().GetUserIdentity(ctx)
	)

	values, err := dao.SysAuthRoles.Ctx(ctx).Fields("DISTINCT "+cols.PlatformCode).Where(cols.SysUserId, identity.Id).
		Cache(gdb.CacheOption{
			Duration: 1 * gtime.D,
			Force:    false,
		}).Array()
	if err != nil {
		return
	}
	if err = dao.SysPlatform.Ctx(ctx).WhereIn(dao.SysAuthRoles.Columns().PlatformCode, values).
		Where(dao.SysPlatform.Columns().Status, 1).
		Cache(gdb.CacheOption{
			Duration: 1 * gtime.D,
			Force:    false,
		}).Scan(&records); err != nil {
		return
	}
	return
}

func (s *sPlatForm) ProxyPath(ctx context.Context, path, method string) (record entity.SysPlatform, err error) {
	err = dao.AuthApi.Ctx(ctx).As("a").Fields("p.*").
		LeftJoin(dao.SysPlatform.Table(), "p", "a.platform_code = p.platform_code").
		Where("a.path = ?", path).Where("a.method = ?", method).Cache(gdb.CacheOption{
		Duration: 1 * gtime.D,
		Force:    false,
	}).Scan(&record)
	return
}

func (s *sPlatForm) verify(ctx context.Context, id int, scoreMap g.Map) (err error) {
	var (
		cols   = dao.SysPlatform.Columns()
		msgMap = g.MapStrStr{
			cols.PlatformCode: "平台编码已存在，请换一个",
			cols.ServerUrl:    "平台地址已存在，请换一个",
		}
	)

	for k, v := range scoreMap {
		count, err := dao.SysPlatform.Ctx(ctx).WhereNot(cols.Id, id).Where(k, v).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return gerror.New(msgMap[k])
		}
	}
	return
}
