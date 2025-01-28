package model

import (
	"gorm.io/gorm"
)

type Products []Product
type Product struct {
	gorm.Model
	CreatedByID uint `gorm:"default:null"`
	UpdatedByID uint `gorm:"default:null"`
	DeletedByID uint `gorm:"default:null"`

	// Required Fields
	Name        string `gorm:"type:varchar(255);not null"`
	Type        string `gorm:"type:text"`
	CompanyCode uint   `gorm:"not null"`
	Description string `gorm:"type:text"`
	Images      string `gorm:"type:text"`
	Available   bool   `gorm:"default:true"`
}
