package schemas

import (
	"time"

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

type TagResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"`
	Code      uint64    `json:"code"`
	Tag       string    `json:"tag"`
	Date      string    `json:"date"`
	Hour      string    `json:"hour"`
	Value     float64   `json:"value"`
}
