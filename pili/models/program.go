package models

import (
	"github.com/daiguadaidai/poow/utils/types"
)

type Program struct {
	Model
	Title        types.NullString `json:"title" gorm:"column:title;not null;default:''"`
	FileName     types.NullString `json:"file_name" gorm:"column:file_name;not null;default:''"`
	HaveDedicate bool             `json:"have_dedicate" gorm:"column:have_dedicate;not null;default:0"`
	Params       types.NullString `json:"params" gorm:"column:params;not null;default:''"`
}

func (Program) TableName() string {
	return "programs"
}
