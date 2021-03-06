// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// QueueTaskLog is the golang structure for table queue_task_log.
type QueueTaskLog struct {
	Id          uint        `json:"id"            ` // ID
	QueueTaskId int         `json:"queue_task_id" ` // 队列任务id
	Log         string      `json:"log"           ` // 错误日志
	FailAmount  int         `json:"fail_amount"   ` // 第几次失败
	CreateAt    *gtime.Time `json:"create_at"     ` // 创建时间
	UpdateAt    *gtime.Time `json:"update_at"     ` // 更新时间
}
