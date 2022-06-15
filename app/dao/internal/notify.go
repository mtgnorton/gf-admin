// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NotifyDao is the data access object for table notify.
type NotifyDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns NotifyColumns // columns contains all the column names of Table for convenient usage.
}

// NotifyColumns defines and stores column names for table notify.
type NotifyColumns struct {
	Id                 string // ID
	Type               string // 交易类型, recharge 充值，withdraw 提现
	RelationId         string // 关联表id
	NotifyData         string // 通知数据
	NotifyAddress      string // 通知地址
	UniqueId           string // 唯一id
	FailAmount         string // 失败次数
	Status             string // 状态 fail 失败,wait等待通知  finish通知完成
	IsImmediatelyRetry string // 是否立即重试
	CreateAt           string // 创建时间
	NotifyAt           string // 上次通知时间
}

//  notifyColumns holds the columns for table notify.
var notifyColumns = NotifyColumns{
	Id:                 "id",
	Type:               "type",
	RelationId:         "relation_id",
	NotifyData:         "notify_data",
	NotifyAddress:      "notify_address",
	UniqueId:           "unique_id",
	FailAmount:         "fail_amount",
	Status:             "status",
	IsImmediatelyRetry: "is_immediately_retry",
	CreateAt:           "create_at",
	NotifyAt:           "notify_at",
}

// NewNotifyDao creates and returns a new DAO object for table data access.
func NewNotifyDao() *NotifyDao {
	return &NotifyDao{
		group:   "default",
		table:   "notify",
		columns: notifyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NotifyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NotifyDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NotifyDao) Columns() NotifyColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NotifyDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NotifyDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NotifyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
