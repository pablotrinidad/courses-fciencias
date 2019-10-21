package main

import (
	"fmt"
	"github.com/pablotrinidad/courses-fciencias/crawler/pages"
)

func main() {
	for _, major := range pages.FetchAllMajors() {
		fmt.Println(major.Name)
		courses, programs := pages.FetchMajorCourses(major.ExternalID)
		break
		for _, program := range programs {
			fmt.Printf("  + %s: (%s)\n", program.Name, program.GetURL())
			fmt.Printf("\tCourses: %d", len(courses[program.ExternalID]))
			fmt.Printf("\tYear: %d\n", program.Year)
		}
	}
}
