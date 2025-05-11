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
	Id         int    `json:"id"         dc:""`
	NickName   string `json:"nickname"   dc:"昵称"`
	UserName   string `json:"username"   dc:"英文名"`
	Password   string `json:"password"   dc:"密码"`
	Name       string `json:"name"       dc:"中文名"`
	Department string `json:"department" dc:"部门"`
	UserId     string `json:"userId"     dc:"userId"`
	Email      string `json:"email"      dc:"用户邮箱"`
	Phone      string `json:"phone"      dc:"电话"`
	Status     int    `json:"status"     dc:"状态 1 - 开启 2 - 关闭"`
	AvatarUrl  string `json:"avatarUrl"  dc:"头像"`
	Remark     string `json:"remark"     dc:"备注"`
	RoleIds    []int  `json:"roleIds"    dc:"规则id"`
}

type UserResetPwd struct {
	Id       int    `json:"id" v:"required"`
	Password string `json:"password" v:"password#密码长度需在6~18之间" dc:"新密码"`
}
