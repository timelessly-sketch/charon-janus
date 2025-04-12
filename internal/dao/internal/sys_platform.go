// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysPlatformDao is the data access object for the table sys_platform.
type SysPlatformDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysPlatformColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysPlatformColumns defines and stores column names for the table sys_platform.
type SysPlatformColumns struct {
	Id           string //
	PlatformName string // 平台中文名
	PlatformCode string // 平台英文编码
	ServerUrl    string // 服务路由前缀
	DefaultRoute string // 默认路由
	Icon         string // 图标
	Status       string // 状态 1 - 开启 2 - 关闭
	PlatformSort string // 排序
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
	DeletedAt    string // 删除时间
}

// sysPlatformColumns holds the columns for the table sys_platform.
var sysPlatformColumns = SysPlatformColumns{
	Id:           "id",
	PlatformName: "platform_name",
	PlatformCode: "platform_code",
	ServerUrl:    "server_url",
	DefaultRoute: "default_route",
	Icon:         "icon",
	Status:       "status",
	PlatformSort: "platform_sort",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
	DeletedAt:    "deleted_at",
}

// NewSysPlatformDao creates and returns a new DAO object for table data access.
func NewSysPlatformDao(handlers ...gdb.ModelHandler) *SysPlatformDao {
	return &SysPlatformDao{
		group:    "default",
		table:    "sys_platform",
		columns:  sysPlatformColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysPlatformDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysPlatformDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysPlatformDao) Columns() SysPlatformColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysPlatformDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysPlatformDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysPlatformDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
