package system

import (
	"charon-janus/internal/dao"
	"charon-janus/internal/library/token"
	"charon-janus/internal/model"
	"charon-janus/internal/model/entity"
	"charon-janus/internal/model/input"
	"charon-janus/internal/service"
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sLogin struct{}

func NewLogin() *sLogin {
	return &sLogin{}
}

func init() {
	service.RegisterLogin(NewLogin())
}

func (s *sLogin) Login(ctx context.Context, inp *input.AccountLoginInp) (records input.LoginModel, err error) {
	var (
		ipa  string
		user entity.SysUser
	)
	if err = dao.SysUser.Ctx(ctx).Where(dao.SysUser.Columns().NickName, inp.Nickname).Scan(&user); err != nil || gerror.Is(err, sql.ErrNoRows) {
		g.Log().Warning(ctx, err)
		return records, gerror.New("用户不存在")
	}

	if inp.FreeIpa {
		ipa = ""
	}
	if s.generatePassword(inp.Password, ipa) != user.Password {
		return records, gerror.New("密码错误")
	}
	generateJWT, err := token.GenerateJWT(ctx, &model.Identity{
		Id:       user.Id,
		Nickname: user.NickName,
		Username: user.UserName,
		Name:     user.Name,
		UserId:   user.UserId,
	})
	if err != nil {
		g.Log().Warning(ctx, err)
		return
	}
	records = input.LoginModel{
		Id:       user.Id,
		Avatar:   user.AvatarUrl,
		Username: user.UserName,
		Nickname: user.NickName,
		Name:     user.Name,
		Token:    generateJWT,
		Role:     []string{"super"},
	}

	return
}

func (s *sLogin) generatePassword(password, ipa string) string {
	hash := sha256.New()
	hash.Write(gconv.Bytes(password))
	return hex.EncodeToString(hash.Sum(nil)) + ipa
}
