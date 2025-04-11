// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPlatform is the golang structure of table sys_platform for DAO operations like Where/Data.
type SysPlatform struct {
	g.Meta       `orm:"table:sys_platform, do:true"`
	Id           interface{} //
	PlatformName interface{} // 平台中文名
	PlatformCode interface{} // 平台英文编码
	ServerUrl    interface{} // 服务路由前缀
	DefaultRoute interface{} // 默认路由
	Icon         interface{} // 图标
	Status       interface{} // 状态 1 - 开启 2 - 关闭
	Sort         interface{} // 排序
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
	DeletedAt    *gtime.Time // 删除时间
}
