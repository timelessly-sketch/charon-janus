// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysAuthRoles is the golang structure for table sys_auth_roles.
type SysAuthRoles struct {
	SysUserId    int         `json:"sysUserId"    orm:"sys_user_id"   description:"用户的id"` // 用户的id
	AuthRoleId   int         `json:"authRoleId"   orm:"auth_role_id"  description:"权限id"`  // 权限id
	PlatformCode string      `json:"platformCode" orm:"platform_code" description:"平台编码"`  // 平台编码
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`  // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"创建时间"`  // 创建时间
}
