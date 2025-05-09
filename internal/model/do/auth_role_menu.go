// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthRoleMenu is the golang structure of table auth_role_menu for DAO operations like Where/Data.
type AuthRoleMenu struct {
	g.Meta    `orm:"table:auth_role_menu, do:true"`
	RoleId    interface{} // 权限ID
	MenuId    interface{} // 菜单ID
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
