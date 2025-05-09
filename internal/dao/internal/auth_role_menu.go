// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthRoleMenuDao is the data access object for the table auth_role_menu.
type AuthRoleMenuDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  AuthRoleMenuColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// AuthRoleMenuColumns defines and stores column names for the table auth_role_menu.
type AuthRoleMenuColumns struct {
	RoleId    string // 权限ID
	MenuId    string // 菜单ID
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// authRoleMenuColumns holds the columns for the table auth_role_menu.
var authRoleMenuColumns = AuthRoleMenuColumns{
	RoleId:    "role_id",
	MenuId:    "menu_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewAuthRoleMenuDao creates and returns a new DAO object for table data access.
func NewAuthRoleMenuDao(handlers ...gdb.ModelHandler) *AuthRoleMenuDao {
	return &AuthRoleMenuDao{
		group:    "default",
		table:    "auth_role_menu",
		columns:  authRoleMenuColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AuthRoleMenuDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AuthRoleMenuDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AuthRoleMenuDao) Columns() AuthRoleMenuColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AuthRoleMenuDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AuthRoleMenuDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AuthRoleMenuDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
