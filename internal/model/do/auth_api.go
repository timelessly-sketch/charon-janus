// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthApi is the golang structure of table auth_api for DAO operations like Where/Data.
type AuthApi struct {
	g.Meta       `orm:"table:auth_api, do:true"`
	Id           interface{} // 接口ID
	Pid          interface{} // 接口PID
	Name         interface{} // 接口名
	Icon         interface{} // 图标
	Title        interface{} // 标题
	Path         interface{} // 接口路径
	Method       interface{} // 接口方法-目录为空,接口不能为空
	ApiType      interface{} // 接口或者目录
	PlatformCode interface{} // 平台标识
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 删除时间
}
