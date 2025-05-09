// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthRoleApiDao is the data access object for the table auth_role_api.
type AuthRoleApiDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AuthRoleApiColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AuthRoleApiColumns defines and stores column names for the table auth_role_api.
type AuthRoleApiColumns struct {
	RoleId    string // 规则id
	ApiId     string // 接口id
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string // 删除时间
}

// authRoleApiColumns holds the columns for the table auth_role_api.
var authRoleApiColumns = AuthRoleApiColumns{
	RoleId:    "role_id",
	ApiId:     "api_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewAuthRoleApiDao creates and returns a new DAO object for table data access.
func NewAuthRoleApiDao(handlers ...gdb.ModelHandler) *AuthRoleApiDao {
	return &AuthRoleApiDao{
		group:    "default",
		table:    "auth_role_api",
		columns:  authRoleApiColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AuthRoleApiDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AuthRoleApiDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AuthRoleApiDao) Columns() AuthRoleApiColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AuthRoleApiDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AuthRoleApiDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AuthRoleApiDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
