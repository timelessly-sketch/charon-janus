// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthRoleApi is the golang structure of table auth_role_api for DAO operations like Where/Data.
type AuthRoleApi struct {
	g.Meta    `orm:"table:auth_role_api, do:true"`
	RoleId    interface{} // 规则id
	ApiId     interface{} // 接口id
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 更新时间
	DeletedAt *gtime.Time // 删除时间
}
