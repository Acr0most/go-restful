package model

import (
	"gorm.io/gorm"
	"time"
)

type CommonModelFields struct {
	ID        uint           `gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at" time_format:"sql_datetime" time_location:"UTC"`
	UpdatedAt *time.Time     `json:"updated_at" time_format:"sql_datetime" time_location:"UTC"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" time_format:"sql_datetime" time_location:"UTC" gorm:"index"`
}
