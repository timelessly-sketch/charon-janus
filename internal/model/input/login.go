package input

import "charon-janus/internal/model/entity"

type AccountLoginInp struct {
	UserName string `json:"username" v:"required#用户名不能为空" dc:"用户名"`
	Password string `json:"password" v:"required#密码不能为空" dc:"密码"`
	FreeIpa  bool   `json:"freeIpa"`
}

type LoginModel struct {
	Id       int      `json:"id" dc:"用户ID"`
	Avatar   string   `json:"avatar" dc:"头像"`
	Username string   `json:"username" dc:"用户名"`
	Nickname string   `json:"nickname" dc:"昵称"`
	Name     string   `json:"name" dc:"中文名"`
	Token    string   `json:"token" dc:"token"`
	Role     []string `json:"role" dc:"权限组"`
}

type UserRoutes struct {
	entity.AuthMenu
}

type LoginLogInp struct {
	Username string `json:"username"`
	PageReq
}

type LoginLogList struct {
	entity.SysLoginLog
}
