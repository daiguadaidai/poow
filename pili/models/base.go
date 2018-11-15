package models

import "github.com/daiguadaidai/poow/utils/types"

// model的级别字段, 每个model都有 ...
type Model struct {
	ID        types.NullInt64 `json:"id"`
	UpdatedAt types.NullTime  `json:"updated_at" gorm:"column:updated_at`
	CreatedAt types.NullTime  `json:"created_at" gorm:"column:created_at`
}
