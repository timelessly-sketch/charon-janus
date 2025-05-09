package user

import (
	"charon-janus/internal/model/input"
	"github.com/gogf/gf/v2/frame/g"
)

type ListReq struct {
	g.Meta `path:"/system/user/list" method:"get" summary:"获取用户列表"`
	input.UserInput
}

type ListRes struct {
	Records []input.UserModelList `json:"records"`
	Total   int                   `json:"total"`
}

type DetailReq struct {
	g.Meta `path:"/system/user/detail/:id" method:"get" summary:"获取用户详情"`
	Id     int `json:"id" v:"required#用户ID缺失"`
}
type DetailRes struct {
	input.UserModelDetail
}

type EditReq struct {
	g.Meta `path:"/system/user/edit" method:"post" summary:"编辑用户"`
	input.UserEditInput
}
type EditRes struct{}

type ResetReq struct {
	g.Meta   `path:"/system/user/reset" method:"post" summary:"重置密码"`
	Username string `json:"Username" v:"required#用户名缺失"`
}

type ResetRes struct{}
