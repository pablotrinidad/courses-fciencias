package entities

import (
	"fmt"
	"github.com/pablotrinidad/courses-fciencias/crawler"
)

const (
	MajorsActuary            = 101
	MajorsBiology            = 201
	MajorsComputerScience    = 104
	MajorsGeoScience         = 127
	MajorsPhysics            = 106
	MajorsMedicalPhysics     = 134
	MajorsMathematics        = 122
	MajorsAppliedMathematics = 136
	MajorsURL                = "licenciatura/resumen"
)

var Majors = []int{
	MajorsActuary, MajorsBiology, MajorsComputerScience, MajorsGeoScience,
	MajorsPhysics, MajorsMedicalPhysics, MajorsMathematics, MajorsAppliedMathematics}

type Major struct {
	BaseEntity
	Name string `json:"name"`
}

// GetURL of the major's main page
func (p *Major) GetURL() string {
	return fmt.Sprintf("%s/%s/%d", crawler.BaseURL, MajorsURL, p.ExternalID)
}
