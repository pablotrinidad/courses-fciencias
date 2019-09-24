package pages

import (
	"github.com/pablotrinidad/courses-fciencias/crawler"
	"github.com/pablotrinidad/courses-fciencias/crawler/entities"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func FetchProgramCourses(program *entities.Program) (data []*entities.ProgramCourse) {
	document := crawler.GetDocument(program.GetURL())

	parenthesisRegex := regexp.MustCompile(`\(([^\)]+)\)`)
	digitRegex := regexp.MustCompile(`\d+`)

	programRawName := document.Find("h1").First().Text()
	programRawName = parenthesisRegex.FindString(programRawName)
	program.Name = strings.Title(programRawName[1 : len(programRawName)-1])

	program.Year, _ = strconv.Atoi(digitRegex.FindString(program.Name))
	return data
}

// FetchMajorCourses
func FetchMajorCourses(major int) (map[int][]*entities.ProgramCourse, map[int]*entities.Program) {
	courses := make(map[int][]*entities.ProgramCourse)
	programs := make(map[int]*entities.Program)

	var wg sync.WaitGroup
	cn := make(chan *entities.ProgramCourse)

	for _, programID := range entities.Programs[major] {
		wg.Add(1)
		programs[programID] = &entities.Program{
			BaseEntity: entities.BaseEntity{ExternalID: programID},
			Major:      major,
			Name:       "",
			Year:       0,
		}
		go func(id int, cn chan *entities.ProgramCourse, program *entities.Program) {
			defer wg.Done()
			for _, course := range FetchProgramCourses(program) {
				cn <- course
			}
		}(programID, cn, programs[programID])
	}

	go func() {
		wg.Wait()
		close(cn)
	}()
	for course := range cn {
		courses[course.Program] = append(courses[course.Program], course)
	}

	return courses, programs
}
