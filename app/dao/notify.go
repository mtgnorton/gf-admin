// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gf-admin/app/dao/internal"
)

// internalNotifyDao is internal type for wrapping internal DAO implements.
type internalNotifyDao = *internal.NotifyDao

// notifyDao is the data access object for table notify.
// You can define custom methods on it to extend its functionality as you wish.
type notifyDao struct {
	internalNotifyDao
}

var (
	// Notify is globally public accessible object for table notify operations.
	Notify = notifyDao{
		internal.NewNotifyDao(),
	}
)

// Fill with you ideas below.
