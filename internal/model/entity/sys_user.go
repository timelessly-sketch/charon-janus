// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure for table sys_user.
type SysUser struct {
	Id         int         `json:"id"         orm:"id"         description:""`                 //
	NickName   string      `json:"nickName"   orm:"nick_name"  description:"昵称"`               // 昵称
	UserName   string      `json:"userName"   orm:"user_name"  description:"英文名"`              // 英文名
	Password   string      `json:"password"   orm:"password"   description:"密码"`               // 密码
	Name       string      `json:"name"       orm:"name"       description:"中文名"`              // 中文名
	Department string      `json:"department" orm:"department" description:"部门"`               // 部门
	UserId     string      `json:"userId"     orm:"user_id"    description:"userId"`           // userId
	Email      string      `json:"email"      orm:"email"      description:"用户邮箱"`             // 用户邮箱
	Phone      string      `json:"phone"      orm:"phone"      description:"电话"`               // 电话
	Status     int         `json:"status"     orm:"status"     description:"状态 1 - 开启 2 - 关闭"` // 状态 1 - 开启 2 - 关闭
	AvatarUrl  string      `json:"avatarUrl"  orm:"avatar_url" description:"头像"`               // 头像
	Remark     string      `json:"remark"     orm:"remark"     description:"备注"`               // 备注
	CreatedAt  *gtime.Time `json:"createdAt"  orm:"created_at" description:"创建时间"`             // 创建时间
	UpdatedAt  *gtime.Time `json:"updatedAt"  orm:"updated_at" description:"更新时间"`             // 更新时间
	DeletedAt  *gtime.Time `json:"deletedAt"  orm:"deleted_at" description:"删除时间"`             // 删除时间
}
