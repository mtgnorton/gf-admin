// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Withdraws is the golang structure for table withdraws.
type Withdraws struct {
	Id              uint        `json:"id"                ` // ID
	UserId          int         `json:"user_id"           ` // 内部用户id
	UserAddress     string      `json:"user_address"      ` // 用户地址
	ExternalOrderId string      `json:"external_order_id" ` // 外部订单id
	ExternalUserId  string      `json:"external_user_id"  ` // 外部用户id
	Hash            string      `json:"hash"              ` // hash
	Symbol          string      `json:"symbol"            ` // 代币符号
	ContractAddress string      `json:"contract_address"  ` //
	From            string      `json:"from"              ` // 转出地址
	To              string      `json:"to"                ` // 转入地址
	Value           string      `json:"value"             ` // 转出金额
	Status          string      `json:"status"            ` // 状态 fail 失败，wait 待转出，process 转出中，wait_notify转出完成,finish_notify通知完成
	CreateAt        *gtime.Time `json:"create_at"         ` // 创建时间
	UpdateAt        *gtime.Time `json:"update_at"         ` // 更新时间
}
