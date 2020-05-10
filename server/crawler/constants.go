package crawler

import (
	"time"
)

const (
	// Website constatns
	baseURL     = "http://www.fciencias.unam.mx"
	majorsURL   = "licenciatura/resumen"
	programsURL = "licenciatura/mapa"
	coursesURL  = "licenciatura/asignaturas"

	// Crawler parameters
	httpUserAgent  = "CoursesCrawlerBot v.1.0 https://github.com/pablotrinidad/courses-fciencias | Download courses catalog"
	requestTimeout = time.Duration(2 * time.Second)

	// Majors
	majorsActuary            = 101
	majorsBiology            = 201
	majorsComputerScience    = 104
	majorsGeoScience         = 127
	majorsPhysics            = 106
	majorsMedicalPhysics     = 134
	majorsMathematics        = 122
	majorsAppliedMathematics = 136
)

var (
	programs = map[int][]int{
		majorsActuary:            []int{2017, 1176, 119, 214},
		majorsBiology:            []int{181, 215},
		majorsComputerScience:    []int{1556, 218},
		majorsGeoScience:         []int{1439, 1440, 1441, 1442, 1443, 1444},
		majorsPhysics:            []int{1081, 216},
		majorsMedicalPhysics:     []int{2016, 4027, 4028, 4138},
		majorsMathematics:        []int{217, 839},
		majorsAppliedMathematics: []int{2055},
	}
	majorsList = []int{
		majorsActuary, majorsBiology, majorsComputerScience, majorsGeoScience,
		majorsPhysics, majorsMedicalPhysics, majorsMathematics, majorsAppliedMathematics,
	}
)
