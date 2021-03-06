// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NotifyLogDao is the data access object for table notify_log.
type NotifyLogDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns NotifyLogColumns // columns contains all the column names of Table for convenient usage.
}

// NotifyLogColumns defines and stores column names for table notify_log.
type NotifyLogColumns struct {
	Id         string // ID
	NotifyId   string // notify id
	Log        string // 错误日志
	FailAmount string // 第几次失败
	CreateAt   string // 创建时间
	UpdateAt   string // 更新时间
}

//  notifyLogColumns holds the columns for table notify_log.
var notifyLogColumns = NotifyLogColumns{
	Id:         "id",
	NotifyId:   "notify_id",
	Log:        "log",
	FailAmount: "fail_amount",
	CreateAt:   "create_at",
	UpdateAt:   "update_at",
}

// NewNotifyLogDao creates and returns a new DAO object for table data access.
func NewNotifyLogDao() *NotifyLogDao {
	return &NotifyLogDao{
		group:   "default",
		table:   "notify_log",
		columns: notifyLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NotifyLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NotifyLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NotifyLogDao) Columns() NotifyLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NotifyLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NotifyLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NotifyLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
