package login

import (
	"charon-janus/internal/model/input"
	"github.com/gogf/gf/v2/frame/g"
)

type AccountLoginReq struct {
	g.Meta `path:"/auth/login" method:"POST" dc:"登录" noAuth:"true"`
	input.AccountLoginInp
}

type AccountLoginRes struct {
	input.LoginModel
}

type RoutesReq struct {
	g.Meta       `path:"/auth/routes" method:"GET" dc:"获取用户权限"`
	PlatFormCode string `json:"platformCode"`
}

type RoutesRes struct {
	Records []input.UserRoutes `json:"records"`
}

type ListReq struct {
	g.Meta `path:"/system/user/loginLog" method:"GET" dc:"登录情况"`
	input.LoginLogInp
}

type ListRes struct {
	Records []input.LoginLogList `json:"records"`
	Total   int                  `json:"total"`
}
