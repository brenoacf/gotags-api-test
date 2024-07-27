package schemas

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Code  uint64
	Tag   string
	Date  string
	Hour  string
	Value float64
}
