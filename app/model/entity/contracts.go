// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Contracts is the golang structure for table contracts.
type Contracts struct {
	Id            int         `json:"id"              ` //
	Symbol        string      `json:"symbol"          ` // 货币类型
	Address       string      `json:"address"         ` // 合约地址
	Decimals      int         `json:"decimals"        ` // 小数位数
	IsCollectOpen int         `json:"is_collect_open" ` // 是否开启,1是 0否
	CreateAt      *gtime.Time `json:"create_at"       ` // 创建时间
	UpdateAt      *gtime.Time `json:"update_at"       ` // 更新时间
}
