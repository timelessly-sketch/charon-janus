// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAuthRoles is the golang structure of table sys_auth_roles for DAO operations like Where/Data.
type SysAuthRoles struct {
	g.Meta     `orm:"table:sys_auth_roles, do:true"`
	SysUserId  interface{} // 用户的id
	AuthRoleId interface{} // 权限id
	CreatedAt  *gtime.Time // 创建时间
	UpdatedAt  *gtime.Time // 创建时间
	DeletedAt  *gtime.Time // 删除时间
}
