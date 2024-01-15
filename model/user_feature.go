package model

import (
	"time"

	"github.com/lib/pq"
)

type UserFeature struct {
	UserID    string          `gorm:"primary_key;not null" json:"user_id"`
	Features  pq.Float32Array `gorm:"type:double precision[]" json:"features"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserFeature(userID string, features []float32) *UserFeature {
	return &UserFeature{
		UserID:   userID,
		Features: features,
	}
}
