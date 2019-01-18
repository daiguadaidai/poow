package models

import (
	"github.com/daiguadaidai/poow/utils/types"
)

type Host struct {
	Model
	IsValid          bool             `json:"is_valid" gorm:"column:is_valid"`
	IsDedicate       bool             `json:"is_dedicate" gorm:"column:is_dedicate"`
	Host             types.NullString `json:"host" gorm:"column:host"`
	RunningTaskCount types.NullInt64  `json:"running_task_count" gorm:"column:running_task_count"`
}

func (Host) TableName() string {
	return "hosts"
}
