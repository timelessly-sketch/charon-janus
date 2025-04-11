// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AuthRoleDao is the data access object for the table auth_role.
type AuthRoleDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AuthRoleColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AuthRoleColumns defines and stores column names for the table auth_role.
type AuthRoleColumns struct {
	Id           string //
	RoleName     string // 角色名
	RoleKey      string // 角色编码
	RoleSort     string // 排序
	PlatformCode string // 平台编码
	Status       string // 状态 1 - 开启 2 - 关闭
	CreatedBy    string // 创建者
	CreatedAt    string // 删除时间
	UpdatedAt    string // 更新时间
	DeletedAt    string // 删除时间
}

// authRoleColumns holds the columns for the table auth_role.
var authRoleColumns = AuthRoleColumns{
	Id:           "id",
	RoleName:     "role_name",
	RoleKey:      "role_key",
	RoleSort:     "role_sort",
	PlatformCode: "platform_code",
	Status:       "status",
	CreatedBy:    "created_by",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// NewAuthRoleDao creates and returns a new DAO object for table data access.
func NewAuthRoleDao(handlers ...gdb.ModelHandler) *AuthRoleDao {
	return &AuthRoleDao{
		group:    "default",
		table:    "auth_role",
		columns:  authRoleColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AuthRoleDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AuthRoleDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AuthRoleDao) Columns() AuthRoleColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AuthRoleDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AuthRoleDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AuthRoleDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
