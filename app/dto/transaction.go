// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package dto

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Transaction is the golang structure of table transaction for DAO operations like Where/Data.
type Transaction struct {
	g.Meta       `orm:"table:transaction, do:true"`
	Id           interface{} // ID
	Symbol       interface{} //
	Hash         interface{} // 交易hash
	Nonce        interface{} // 交易nonce
	From         interface{} // 转出地址
	To           interface{} // 转入地址
	Value        interface{} // 转账金额,单位为wei
	ValueDecimal interface{} // 转账金额,小数
	GasUsed      interface{} // gas 实际消耗
	GasPrice     interface{} // gas 价格
	BlockNumber  interface{} // 块号
	BlockHash    interface{} // 块hash
	Type         interface{} // 转账类型
	CreateAt     *gtime.Time // 创建时间
	UpdateAt     *gtime.Time // 更新时间
}