// fciencias enty point

package fciencias

// FetchAllData initiate the data download from the website
func FetchAllData() *[]Major {
	majors := FetchMajors()
	return &majors
}
