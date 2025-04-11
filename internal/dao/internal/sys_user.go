// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SysUserDao is the data access object for the table sys_user.
type SysUserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SysUserColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SysUserColumns defines and stores column names for the table sys_user.
type SysUserColumns struct {
	Id         string //
	NickName   string // 昵称
	UserName   string // 英文名
	Password   string // 密码
	Name       string // 中文名
	Department string // 部门
	UserId     string // userId
	Email      string // 用户邮箱
	Phone      string // 电话
	Status     string // 状态 1 - 开启 2 - 关闭
	AvatarUrl  string // 头像
	Remark     string // 备注
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
	DeletedAt  string // 删除时间
}

// sysUserColumns holds the columns for the table sys_user.
var sysUserColumns = SysUserColumns{
	Id:         "id",
	NickName:   "nick_name",
	UserName:   "user_name",
	Password:   "password",
	Name:       "name",
	Department: "department",
	UserId:     "user_id",
	Email:      "email",
	Phone:      "phone",
	Status:     "status",
	AvatarUrl:  "avatar_url",
	Remark:     "remark",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewSysUserDao creates and returns a new DAO object for table data access.
func NewSysUserDao(handlers ...gdb.ModelHandler) *SysUserDao {
	return &SysUserDao{
		group:    "default",
		table:    "sys_user",
		columns:  sysUserColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SysUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SysUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SysUserDao) Columns() SysUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SysUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SysUserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SysUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
