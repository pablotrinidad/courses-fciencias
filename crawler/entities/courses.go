package entities

const CoursesURL = "http://www.fciencias.unam.mx/licenciatura/asignaturas/"

type Course struct {
	BaseEntity
	Name string `json:"name"`
}

type ProgramCourse struct {
	BaseEntity
	Program   int    `json:"program"`
	Name      string `json:"name"`
	Semester  int    `json:"semester"`
	Credits   int    `json:"credits"`
	Syllabus  string `json:"syllabus"`
	Mandatory bool   `json:"mandatory"`
}
