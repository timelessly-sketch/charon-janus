// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthRole is the golang structure for table auth_role.
type AuthRole struct {
	Id           int         `json:"id"           orm:"id"            description:""`                 //
	RoleName     string      `json:"roleName"     orm:"role_name"     description:"角色名"`              // 角色名
	RoleKey      string      `json:"roleKey"      orm:"role_key"      description:"角色编码"`             // 角色编码
	RoleSort     int         `json:"roleSort"     orm:"role_sort"     description:"排序"`               // 排序
	PlatformCode string      `json:"platformCode" orm:"platform_code" description:"平台编码"`             // 平台编码
	Status       int         `json:"status"       orm:"status"        description:"状态 1 - 开启 2 - 关闭"` // 状态 1 - 开启 2 - 关闭
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"删除时间"`             // 删除时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`             // 更新时间
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:"删除时间"`             // 删除时间
}
