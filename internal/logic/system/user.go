package system

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
)

type sUser struct{}

func NewUser() *sUser {
	return &sUser{}
}

func init() {
	service.RegisterUser(NewUser())
}

func (s *sUser) List(ctx context.Context, inp *input.UserInput) (records []input.UserModelList, total int, err error) {
	cols := dao.SysUser.Columns()
	db := dao.SysUser.Ctx(ctx)
	if inp.Name != "" {
		db = db.WhereLike(cols.Name, "%"+inp.Name+"%")
	}
	if inp.UserName != "" {
		db = db.WhereLike(cols.UserName, "%"+inp.UserName+"%")
	}
	err = db.Page(inp.Page, inp.Size).ScanAndCount(&records, &total, true)
	return
}

func (s *sUser) Detail(ctx context.Context, id int) (records input.UserModelDetail, err error) {
	var (
		roles []entity.AuthRole
	)
	if err = dao.AuthRole.Ctx(ctx).Where(dao.AuthRole.Columns().Status, 1).Scan(&roles); err != nil {
		return
	}
	array, _ := dao.SysAuthRoles.Ctx(ctx).Fields(dao.SysAuthRoles.Columns().AuthRoleId).
		Where(dao.SysAuthRoles.Columns().SysUserId, id).Array()

	records = input.UserModelDetail{
		RoleIds: gvar.New(array).Ints(),
		Roles:   roles,
	}
	return
}

func (s *sUser) Edit(ctx context.Context, inp *input.UserEditInput) (err error) {
	var (
		cols     = dao.SysUser.Columns()
		authList = make([]entity.SysAuthRoles, 0)
	)

	if err = s.verify(ctx, inp.Id, g.Map{
		cols.NickName: inp.Name,
		cols.UserName: inp.UserName,
		cols.Email:    inp.Email,
		cols.Phone:    inp.Phone,
	}); err != nil {
		return err
	}

	for _, id := range inp.RoleIds {
		authList = append(authList, entity.SysAuthRoles{
			SysUserId:  inp.Id,
			AuthRoleId: id,
		})
	}
	if inp.Id == 0 {
		return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
			if _, err = dao.SysUser.Ctx(ctx).Data(&inp.SysUser).Insert(); err != nil {
				return err
			}
			if len(authList) > 0 {
				if _, err = dao.SysAuthRoles.Ctx(ctx).Data(authList).Insert(); err != nil {
					return err
				}
			}
			return
		})
	}

	return g.DB().Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		if _, err = dao.SysUser.Ctx(ctx).WherePri(inp.Id).Data(&inp.SysUser).Update(); err != nil {
			return
		}
		if _, err = dao.SysAuthRoles.Ctx(ctx).Where(dao.SysAuthRoles.Columns().SysUserId, inp.Id).Delete(); err != nil {
			return
		}
		if len(authList) > 0 {
			if _, err = dao.SysAuthRoles.Ctx(ctx).Data(authList).Insert(); err != nil {
				return err
			}
		}
		return
	})
}

func (s *sUser) verify(ctx context.Context, id int, scoreMap g.Map) (err error) {
	var (
		cols   = dao.SysUser.Columns()
		msgMap = g.MapStrStr{
			cols.NickName: "昵称已存在，请换一个",
			cols.UserName: "英文名已存在，请换一个",
			cols.Email:    "邮箱已存在，请换一个",
			cols.Phone:    "电话已存在，请换一个",
		}
	)

	for k, v := range scoreMap {
		count, err := dao.SysUser.Ctx(ctx).WhereNot(cols.Id, id).Where(k, v).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return gerror.New(msgMap[k])
		}
	}
	return
}
