// F. Ciencias handlers

package crawler

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/pablotrinidad/courses-fciencias/fciencias"
)

// GetAllMajorsHandler return all majors
func FetchAllDataHandler(w http.ResponseWriter, r *http.Request) {
	majors := fciencias.FetchAllData()
	render.JSON(w, r, majors)
}
