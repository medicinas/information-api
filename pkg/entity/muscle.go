package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Muscle struct {
	gorm.Model
	ID                uint              `gorm:"primary_key;auto_increment" json:"id"`
	Name              string            `gorm:"size:255;not null;unique" json:"name"`
	AnatomicRegion    AnatomicRegion    `gorm:"association_foreignkey:ID" json:"anatomicRegion"`
	SubAnatomicRegion SubAnatomicRegion `gorm:"association_foreignkey:ID" json:"subAnatomicRegion"`
	CreatedAt         time.Time         `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt         time.Time         `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
