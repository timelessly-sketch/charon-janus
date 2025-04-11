package login

import (
	"charon-janus/internal/model/input"
	"github.com/gogf/gf/v2/frame/g"
)

type AccountLoginReq struct {
	g.Meta `path:"/login" method:"post" summary:"登录" noAuth:"true"`
	input.AccountLoginInp
}

type AccountLoginRes struct {
	input.LoginModel
}
