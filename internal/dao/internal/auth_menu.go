// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthMenuDao is the data access object for the table auth_menu.
type AuthMenuDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AuthMenuColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AuthMenuColumns defines and stores column names for the table auth_menu.
type AuthMenuColumns struct {
	Id            string // 菜单ID
	Pid           string // 父菜单ID
	Name          string // 菜单名称（英文标识）
	Path          string // 路由路径（全局唯一）
	Title         string // 显示标题
	RequiresAuth  string // 是否需要鉴权
	Icon          string // 图标类名
	MenuType      string // 菜单类型：目录/页面
	ComponentPath string // 组件文件路径
	Hide          string // 是否隐藏菜单
	ActiveMenu    string // 激活显示的菜单路径
	KeepAlive     string // 是否缓存页面
	WithoutTab    string // 当前路由不会被添加到Tab中
	Href          string // 外部链接地址
	Order         string // 菜单排序权重
	PlatformCode  string // 平台编码
	CreatedAt     string // 创建时间
	UpdatedAt     string // 更新时间
	DeletedAt     string // 删除时间
}

// authMenuColumns holds the columns for the table auth_menu.
var authMenuColumns = AuthMenuColumns{
	Id:            "id",
	Pid:           "pid",
	Name:          "name",
	Path:          "path",
	Title:         "title",
	RequiresAuth:  "requires_auth",
	Icon:          "icon",
	MenuType:      "menu_type",
	ComponentPath: "component_path",
	Hide:          "hide",
	ActiveMenu:    "active_menu",
	KeepAlive:     "keep_alive",
	WithoutTab:    "without_tab",
	Href:          "href",
	Order:         "order",
	PlatformCode:  "platform_code",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
}

// NewAuthMenuDao creates and returns a new DAO object for table data access.
func NewAuthMenuDao(handlers ...gdb.ModelHandler) *AuthMenuDao {
	return &AuthMenuDao{
		group:    "default",
		table:    "auth_menu",
		columns:  authMenuColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AuthMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AuthMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AuthMenuDao) Columns() AuthMenuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AuthMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AuthMenuDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *AuthMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
