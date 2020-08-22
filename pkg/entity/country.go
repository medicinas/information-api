package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Country struct {
	gorm.Model
	ID        uint      `gorm:"primary_key;auto_increment" json:"id"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Code      string    `gorm:"size:10;not null;unique" json:"code"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
