// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysAuthRolesDao is the data access object for the table sys_auth_roles.
type SysAuthRolesDao struct {
	table    string              // table is the underlying table name of the DAO.
	group    string              // group is the database configuration group name of the current DAO.
	columns  SysAuthRolesColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler  // handlers for customized model modification.
}

// SysAuthRolesColumns defines and stores column names for the table sys_auth_roles.
type SysAuthRolesColumns struct {
	SysUserId    string // 用户的id
	AuthRoleId   string // 权限id
	PlatformCode string // 平台编码
	CreatedAt    string // 创建时间
	UpdatedAt    string // 创建时间
}

// sysAuthRolesColumns holds the columns for the table sys_auth_roles.
var sysAuthRolesColumns = SysAuthRolesColumns{
	SysUserId:    "sys_user_id",
	AuthRoleId:   "auth_role_id",
	PlatformCode: "platform_code",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewSysAuthRolesDao creates and returns a new DAO object for table data access.
func NewSysAuthRolesDao(handlers ...gdb.ModelHandler) *SysAuthRolesDao {
	return &SysAuthRolesDao{
		group:    "default",
		table:    "sys_auth_roles",
		columns:  sysAuthRolesColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysAuthRolesDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysAuthRolesDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysAuthRolesDao) Columns() SysAuthRolesColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysAuthRolesDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysAuthRolesDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysAuthRolesDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
