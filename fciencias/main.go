// fciencias enty point

package fciencias

import "github.com/pablotrinidad/courses-fciencias/models"

// FetchAllData initiate the data download from the website
func FetchAllData() *[]models.Major {
	majors := FetchMajors()
	return &majors
}
