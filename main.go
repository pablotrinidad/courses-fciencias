package main

import (
	"fmt"

	"github.com/pablotrinidad/courses-fciencias/crawler/pages"
)

func main() {
	for _, major := range pages.FetchAllMajors() {
		fmt.Println(major.Name)
		pages.FetchMajorCourses(major.ExternalID)
	}
}
