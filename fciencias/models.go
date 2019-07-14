// Exposes type strcutes used accross the package

package fciencias

// baseModel contain fields present in every struct
type baseModel struct {
	ID         int `json:"id"`
	ExternalID int `json:"external_id"`
}

// A Major offered in the faculty
type Major struct {
	baseModel
	Name          string         `json:"name"`
	AcademicPlans []AcademicPlan `json:"academic_plans"`
}

// An AcademicPlan that a major have
type AcademicPlan struct {
	baseModel
	Name string `json:"name"`
	Year int    `json:"year"`
}
