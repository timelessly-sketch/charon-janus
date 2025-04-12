// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AuthMenu is the golang structure for table auth_menu.
type AuthMenu struct {
	Id            int         `json:"id"            orm:"id"             description:"菜单ID"`           // 菜单ID
	Pid           int         `json:"pid"           orm:"pid"            description:"父菜单ID"`          // 父菜单ID
	Name          string      `json:"name"          orm:"name"           description:"菜单名称（英文标识）"`     // 菜单名称（英文标识）
	Path          string      `json:"path"          orm:"path"           description:"路由路径（全局唯一）"`     // 路由路径（全局唯一）
	Title         string      `json:"title"         orm:"title"          description:"显示标题"`           // 显示标题
	RequiresAuth  bool        `json:"requiresAuth"  orm:"requires_auth"  description:"是否需要鉴权"`         // 是否需要鉴权
	Icon          string      `json:"icon"          orm:"icon"           description:"图标类名"`           // 图标类名
	MenuType      string      `json:"menuType"      orm:"menu_type"      description:"菜单类型：目录/页面"`     // 菜单类型：目录/页面
	ComponentPath string      `json:"componentPath" orm:"component_path" description:"组件文件路径"`         // 组件文件路径
	Hide          bool        `json:"hide"          orm:"hide"           description:"是否隐藏菜单"`         // 是否隐藏菜单
	ActiveMenu    string      `json:"activeMenu"    orm:"active_menu"    description:"激活显示的菜单路径"`      // 激活显示的菜单路径
	KeepAlive     bool        `json:"keepAlive"     orm:"keep_alive"     description:"是否缓存页面"`         // 是否缓存页面
	WithoutTab    bool        `json:"withoutTab"    orm:"without_tab"    description:"当前路由不会被添加到Tab中"` // 当前路由不会被添加到Tab中
	Href          string      `json:"href"          orm:"href"           description:"外部链接地址"`         // 外部链接地址
	Order         int         `json:"order"         orm:"order"          description:"菜单排序权重"`         // 菜单排序权重
	PlatformCode  string      `json:"platformCode"  orm:"platform_code"  description:"平台编码"`           // 平台编码
	CreatedAt     *gtime.Time `json:"createdAt"     orm:"created_at"     description:"创建时间"`           // 创建时间
	UpdatedAt     *gtime.Time `json:"updatedAt"     orm:"updated_at"     description:"更新时间"`           // 更新时间
	DeletedAt     *gtime.Time `json:"deletedAt"     orm:"deleted_at"     description:"删除时间"`           // 删除时间
}
