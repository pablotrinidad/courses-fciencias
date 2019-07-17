// Exposes type strcutes used accross the package

package fciencias

import (
	"time"
)

// baseModel contain fields present in every struct
type baseEntity struct {
	ID         int       `json:"id"`
	ExternalID int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
}

// A Major offered in the faculty
type Major struct {
	baseEntity
	Name string `json:"name"`
}

// An AcademicProgram that a major have
type AcademicProgram struct {
	baseEntity
	Name string `json:"name"`
	Year int    `json:"year"`
}
