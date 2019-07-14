// Course FCiencias.
// A web crawler made to download and store UNAM's Faculty of Science
// courses schedules.

package main

import (
	"fmt"

	"github.com/pablotrinidad/courses-fciencias/fciencias"
)

func main() {
	majors := fciencias.GetMajors()
	for _, major := range majors {
		fmt.Printf("%d) %s\n", major.ID, major.Name)
		for _, plan := range major.AcademicPlans {
			fmt.Printf(
				"\t%d (ext: %d)\t%s\tyear: %d\n",
				plan.ID, plan.ExternalID, plan.Name, plan.Year)
		}
	}

}
