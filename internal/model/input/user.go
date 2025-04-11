package input

import "charon-janus/internal/model/entity"

type UserInput struct {
	UserName string `json:"username"`
	Name     string `json:"name"`
	PageReq
}

type UserModelList struct {
	entity.SysUser
}

type UserModelDetail struct {
	RoleIds []int             `json:"roleIds"`
	Roles   []entity.AuthRole `json:"roles"`
}

type UserEditInput struct {
	entity.SysUser
	RoleIds []int `json:"roleIds"`
}
