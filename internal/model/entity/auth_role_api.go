// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthRoleApi is the golang structure for table auth_role_api.
type AuthRoleApi struct {
	RoleId    int         `json:"roleId"    orm:"role_id"    description:"规则id"` // 规则id
	ApiId     int         `json:"apiId"     orm:"api_id"     description:"接口id"` // 接口id
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"` // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"` // 更新时间
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:"删除时间"` // 删除时间
}
