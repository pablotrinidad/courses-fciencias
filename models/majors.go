// Majors models

package models

import (
	"github.com/jinzhu/gorm"
)

// Major offerred in the faculty
type Major struct {
	gorm.Model
	ExternalID int `gorm:"unique_index; not null"`
	Name       string
}

// AcademicProgram is a major's set of courses distributed accross semesters
type AcademicProgram struct {
	gorm.Model
	ExternalID int   `gorm:"unique_index; not null"`
	Major      Major `gorm:"foreignkey:MajorID; not null"`
	MajorID    uint
	Name       string
	Year       uint
}
