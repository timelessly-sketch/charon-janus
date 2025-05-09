// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthRole is the golang structure of table auth_role for DAO operations like Where/Data.
type AuthRole struct {
	g.Meta       `orm:"table:auth_role, do:true"`
	Id           interface{} //
	RoleName     interface{} // 角色名
	RoleKey      interface{} // 角色编码
	RoleSort     interface{} // 排序
	PlatformCode interface{} // 平台编码
	Status       interface{} // 状态 1 - 开启 2 - 关闭
	CreatedAt    *gtime.Time // 删除时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 删除时间
}
