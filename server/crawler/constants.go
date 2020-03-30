package crawler

const (
	baseURL       = "http://www.fciencias.unam.mx"
	majorsURL     = "licenciatura/resumen"
	httpUserAgent = "CoursesCrawlerBot v.1.0 https://github.com/pablotrinidad/courses-fciencias | Download courses catalog"
)

const (
	majorsActuary            = 101
	majorsBiology            = 201
	majorsComputerScience    = 104
	majorsGeoScience         = 127
	majorsPhysics            = 106
	majorsMedicalPhysics     = 134
	majorsMathematics        = 122
	majorsAppliedMathematics = 136
)

var majorsList = []int{
	majorsActuary, majorsBiology, majorsComputerScience, majorsGeoScience,
	majorsPhysics, majorsMedicalPhysics, majorsMathematics, majorsAppliedMathematics}
