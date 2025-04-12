// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SysPlatform is the golang structure for table sys_platform.
type SysPlatform struct {
	Id           int         `json:"id"           orm:"id"            description:""`                 //
	PlatformName string      `json:"platformName" orm:"platform_name" description:"平台中文名"`            // 平台中文名
	PlatformCode string      `json:"platformCode" orm:"platform_code" description:"平台英文编码"`           // 平台英文编码
	ServerUrl    string      `json:"serverUrl"    orm:"server_url"    description:"服务路由前缀"`           // 服务路由前缀
	DefaultRoute string      `json:"defaultRoute" orm:"default_route" description:"默认路由"`             // 默认路由
	Icon         string      `json:"icon"         orm:"icon"          description:"图标"`               // 图标
	Status       int         `json:"status"       orm:"status"        description:"状态 1 - 开启 2 - 关闭"` // 状态 1 - 开启 2 - 关闭
	PlatformSort int         `json:"platformSort" orm:"platform_sort" description:"排序"`               // 排序
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`             // 创建时间
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`             // 更新时间
	DeletedAt    *gtime.Time `json:"deletedAt"    orm:"deleted_at"    description:"删除时间"`             // 删除时间
}
