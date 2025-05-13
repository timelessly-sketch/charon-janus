package login

import (
	"charon-janus/internal/dao"
	"charon-janus/internal/library/cache"
	"charon-janus/internal/library/contexts"
	"charon-janus/internal/library/token"
	"charon-janus/internal/model"
	"charon-janus/internal/model/entity"
	"charon-janus/internal/model/input"
	"charon-janus/internal/service"
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

type sLogin struct{}

func NewLogin() *sLogin {
	return &sLogin{}
}

func init() {
	service.RegisterLogin(NewLogin())
}

func (s *sLogin) Login(ctx context.Context, inp *input.AccountLoginInp) (records input.LoginModel, err error) {
	var user entity.SysUser
	if err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().Username, inp.UserName).Scan(&user); err != nil || gerror.Is(err, sql.ErrNoRows) {
		g.Log().Warning(ctx, err)
		return records, gerror.New("用户不存在")
	}

	if inp.FreeIpa {
		//ipa = ""
	}
	records = input.LoginModel{
		Id:       user.Id,
		Avatar:   user.AvatarUrl,
		Username: user.Username,
		Nickname: user.Nickname,
		Name:     user.Name,
	}

	if s.generatePassword(inp.Password, "") != user.Password {
		return records, gerror.New("密码错误")
	}
	records.Token, err = token.GenerateJWT(ctx, &model.Identity{
		Id:       user.Id,
		Nickname: user.Nickname,
		Username: user.Username,
		Name:     user.Name,
		UserId:   user.UserId,
	})
	if err != nil {
		g.Log().Warning(ctx, err)
		return
	}
	array, err := g.DB().Model("sys_auth_roles sa").Fields("r.role_key").
		LeftJoin("auth_role r", "sa.auth_role_id = r.id").Where("sa.sys_user_id = ?", user.Id).Array()
	if err != nil {
		g.Log().Warning(ctx, err)
		return records, gerror.New("获取权限失败")
	}
	contexts.SetUser(ctx, &model.Identity{
		Id:       user.Id,
		Nickname: user.Nickname,
		Username: user.Username,
		Name:     user.Name,
		UserId:   user.UserId,
	})
	records.Role = gconv.Strings(array)
	return
}

func (s *sLogin) UserRoutes(ctx context.Context, code string) (records []input.UserRoutes, err error) {
	var (
		identity         = contexts.Get(ctx).User
		menuList         = make([]entity.AuthMenu, 0)
		menuIDs          = garray.NewIntArray(true)
		collectParentIDs func(pid int)
		cols             = dao.AuthMenu.Columns()
	)
	if code == "" {
		options, _ := service.PlatForm().Options(ctx)
		if len(options) == 0 {
			return records, gerror.New("未授权任何平台")
		}
		code = options[0].PlatformCode
	}
	get, err := cache.Instance().Get(ctx, s.LoginMenuCacheKey(code, identity.Id))
	if err != nil {
		return
	}
	if !get.IsEmpty() {
		err = gconv.Structs(get, &records)
		return
	}
	err = g.DB().Model("sys_user u").
		Fields("m.*").Distinct().
		LeftJoin("sys_auth_roles ar", "u.id = ar.sys_user_id").
		LeftJoin("auth_role r", "ar.auth_role_id = r.id").
		LeftJoin("auth_role_menu rm", "r.id = rm.role_id").
		LeftJoin("auth_menu m", "rm.menu_id = m.id").
		Where("u.id = ? and m.platform_code = ?", identity.Id, code).Scan(&menuList)
	if err != nil {
		return
	}

	collectParentIDs = func(pid int) {
		if pid == 0 || menuIDs.Contains(pid) {
			return
		}
		menuIDs.Append(pid)
		var parent entity.AuthMenu
		_ = dao.AuthMenu.Ctx(ctx).WherePri(pid).Scan(&parent)

		if parent.Pid != 0 {
			collectParentIDs(parent.Pid)
		}
	}
	for _, menu := range menuList {
		menuIDs.Append(menu.Id)
		if menu.Pid != 0 {
			collectParentIDs(menu.Pid)
		}
	}
	if err = dao.AuthMenu.Ctx(ctx).OrderAsc(cols.Order).WhereIn(cols.Id, gvar.New(menuIDs).Ints()).Scan(&records); err != nil {
		return
	}

	err = cache.Instance().Set(ctx, s.LoginMenuCacheKey(code, identity.Id), records, 8*24*time.Hour)
	return
}

func (s *sLogin) LoginMenuCacheKey(code string, id int) (key string) {
	return fmt.Sprintf("%s:%s:%d", "Login_menu", code, id)
}

func (s *sLogin) generatePassword(password, ipa string) string {
	hash := sha256.New()
	hash.Write(gconv.Bytes(password))
	return hex.EncodeToString(hash.Sum(nil)) + ipa
}
