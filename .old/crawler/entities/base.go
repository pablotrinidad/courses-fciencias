package entities

// BaseEntity contain base fields for every other struct
type BaseEntity struct {
	ExternalID int `json:"external_id"`
}
