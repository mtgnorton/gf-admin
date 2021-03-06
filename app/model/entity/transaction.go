// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Transaction is the golang structure for table transaction.
type Transaction struct {
	Id           uint        `json:"id"            ` // ID
	Symbol       string      `json:"symbol"        ` //
	Hash         string      `json:"hash"          ` // 交易hash
	Nonce        int         `json:"nonce"         ` // 交易nonce
	From         string      `json:"from"          ` // 转出地址
	To           string      `json:"to"            ` // 转入地址
	Value        string      `json:"value"         ` // 转账金额,单位为wei
	ValueDecimal float64     `json:"value_decimal" ` // 转账金额,小数
	GasUsed      int         `json:"gas_used"      ` // gas 实际消耗
	GasPrice     string      `json:"gas_price"     ` // gas 价格
	BlockNumber  int         `json:"block_number"  ` // 块号
	BlockHash    string      `json:"block_hash"    ` // 块hash
	Type         int         `json:"type"          ` // 转账类型
	CreateAt     *gtime.Time `json:"create_at"     ` // 创建时间
	UpdateAt     *gtime.Time `json:"update_at"     ` // 更新时间
}
