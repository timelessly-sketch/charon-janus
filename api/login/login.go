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

type RoutesReq struct {
	g.Meta       `path:"/routes" method:"get" summary:"获取用户权限"`
	PlatFormCode string `json:"platformCode"`
}

type RoutesRes struct {
	input.UserRoutes
}
