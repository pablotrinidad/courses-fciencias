// Exposes type strcutes used accross the package

package fciencias

import (
	"time"

	"cloud.google.com/go/datastore"
)

// baseModel contain fields present in every struct
type baseEntity struct {
	ID         *datastore.Key `json:"-" datastore:"__key__"`
	ExternalID int            `json:"id" datastore:"external_id"`
	CreatedAt  time.Time      `json:"created_at" datastore:"created_at"`
}

// A Major offered in the faculty
type Major struct {
	baseEntity
	Name string `json:"name" datastore:"name"`
}

// An AcademicProgram that a major have
type AcademicProgram struct {
	baseEntity
	Name string `json:"name" datastore:"name"`
	Year int    `json:"year" datastore:"year"`
}
