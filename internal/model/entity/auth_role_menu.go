// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthRoleMenu is the golang structure for table auth_role_menu.
type AuthRoleMenu struct {
	RoleId    int         `json:"roleId"    orm:"role_id"    description:"权限ID"` // 权限ID
	MenuId    int         `json:"menuId"    orm:"menu_id"    description:"菜单ID"` // 菜单ID
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"` // 更新时间
}
