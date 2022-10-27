package entity

import (
	"gorm.io/gorm"
)

type ScholarAdmin struct {
	gorm.Model
	Name         string
	Email        string        `gorm:"uniqueIndex"`
	Scholarships []Scholarship `gorm:"foreignKey:ScholarAdminID"`
}

type ScholarStatus struct {
	gorm.Model
	StatusName   string
	Scholarships []Scholarship `gorm:"foreignKey:ScholarStatusID"`
}

type ScholarType struct {
	gorm.Model
	TypeName     string
	Scholarships []Scholarship `gorm:"foreignKey:ScholarTypeID"`
}

type Scholarship struct {
	gorm.Model
	ScholarName    string
	ScholarAdminID *uint
	ScholarAdmin   ScholarAdmin

	ScholarStatusID *uint
	ScholarStatus   ScholarStatus

	ScholarTypeID *uint
	ScholarType   ScholarType

	ScholarDetail string
}
