package models

import (
	"github.com/daiguadaidai/poow/utils/types"
)

const (
	_ = iota
	TASK_STATUS_SUCCESS
	TASK_STATUS_RUNNING
	TASK_STATUS_FAIL
)

type Task struct {
	Model
	ProgramId  types.NullInt64  `json:"program_id "gorm:"column:program_id;not null"`
	TaskUUID   types.NullString `json:"task_uuid" gorm:"column:task_uuid;not null"`
	Host       types.NullString `json:"host" gorm:"column:host;not null"`
	FileName   types.NullString `json:"file_name" gorm:"column:file_name;not null"`
	Params     types.NullString `json:"params" gorm:"column:params;not null;default:''"`
	Pid        types.NullInt64  `json:"pid" gorm:"column:pid;not null;default:0"`
	LogPath    types.NullString `json:"log_path" gorm:"column:log_path;not null;default:''"`
	NotifyInfo types.NullString `json:"notify_info" gorm:"column:notify_info;not null;default:''"`
	RealInfo   types.NullString `json:"real_info" gorm:"column:real_info;not null;default:''"`
	Status     types.NullInt64  `json:"status" gorm:"column:status;not null;default:2"`
}

func (Task) TableName() string {
	return "tasks"
}
