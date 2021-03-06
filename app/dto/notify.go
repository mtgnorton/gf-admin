// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Notify is the golang structure of table notify for DAO operations like Where/Data.
type Notify struct {
	g.Meta             `orm:"table:notify, do:true"`
	Id                 interface{} // ID
	Type               interface{} // 交易类型, recharge 充值，withdraw 提现
	RelationId         interface{} // 关联表id
	NotifyData         interface{} // 通知数据
	NotifyAddress      interface{} // 通知地址
	UniqueId           interface{} // 唯一id
	FailAmount         interface{} // 失败次数
	Status             interface{} // 状态 fail 失败,wait等待通知  finish通知完成
	IsImmediatelyRetry interface{} // 是否立即重试
	CreateAt           *gtime.Time // 创建时间
	NotifyAt           *gtime.Time // 上次通知时间
}
