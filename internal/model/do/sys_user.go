// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysUser is the golang structure of table sys_user for DAO operations like Where/Data.
type SysUser struct {
	g.Meta     `orm:"table:sys_user, do:true"`
	Id         interface{} //
	Nickname   interface{} // 昵称
	Username   interface{} // 英文名
	Password   interface{} // 密码
	Name       interface{} // 中文名
	Department interface{} // 部门
	UserId     interface{} // userId
	Email      interface{} // 用户邮箱
	Phone      interface{} // 电话
	Status     interface{} // 状态 1 - 开启 2 - 关闭
	AvatarUrl  interface{} // 头像
	Remark     interface{} // 备注
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 更新时间
	DeletedAt  *gtime.Time // 删除时间
}
