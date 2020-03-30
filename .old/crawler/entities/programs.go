package entities

import (
	"fmt"

	"github.com/pablotrinidad/courses-fciencias/crawler"
)

const ProgramsURL = "licenciatura/mapa"

var Programs = make(map[int][]int)

func init() {
	Programs[MajorsActuary] = []int{2017, 1176, 119, 214}
	Programs[MajorsBiology] = []int{181, 215}
	Programs[MajorsComputerScience] = []int{1556, 218}
	Programs[MajorsGeoScience] = []int{1439, 1440, 1441, 1442, 1443, 1444}
	Programs[MajorsPhysics] = []int{1081, 216}
	Programs[MajorsMedicalPhysics] = []int{2016, 4027, 4028, 4138}
	Programs[MajorsMathematics] = []int{217, 839}
	Programs[MajorsAppliedMathematics] = []int{2055}
}

type Program struct {
	BaseEntity
	Major int    `json:"major"`
	Name  string `	json:"name"`
	Year  int    `json:"year"`
}

// GetURL of the program's main page
func (p *Program) GetURL() string {
	return fmt.Sprintf("%s/%s/%d/%d", crawler.BaseURL, ProgramsURL, p.Major, p.ExternalID)
}
