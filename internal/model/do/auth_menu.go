// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthMenu is the golang structure of table auth_menu for DAO operations like Where/Data.
type AuthMenu struct {
	g.Meta        `orm:"table:auth_menu, do:true"`
	Id            interface{} // 菜单ID
	Pid           interface{} // 父菜单ID
	Name          interface{} // 菜单名称（英文标识）
	Path          interface{} // 路由路径（全局唯一）
	Title         interface{} // 显示标题
	RequiresAuth  interface{} // 是否需要鉴权
	Icon          interface{} // 图标类名
	MenuType      interface{} // 菜单类型：目录/页面
	ComponentPath interface{} // 组件文件路径
	Hide          interface{} // 是否隐藏菜单
	ActiveMenu    interface{} // 激活显示的菜单路径
	KeepAlive     interface{} // 是否缓存页面
	WithoutTab    interface{} // 当前路由不会被添加到Tab中
	Href          interface{} // 外部链接地址
	Order         interface{} // 菜单排序权重
	PlatformCode  interface{} // 平台编码
	CreatedAt     *gtime.Time // 创建时间
	UpdatedAt     *gtime.Time // 更新时间
	DeletedAt     *gtime.Time // 删除时间
}
