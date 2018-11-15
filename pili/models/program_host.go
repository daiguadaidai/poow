package models

import (
	"github.com/daiguadaidai/poow/utils/types"
)

type ProgramHost struct {
	Model
	ProgramId types.NullInt64 `json:"program_id" gorm:"column:program_id"`
	HostId    types.NullInt64 `json:"host_id" gorm:"column:host_id"`
}

func (ProgramHost) TableName() string {
	return "program_hosts"
}
