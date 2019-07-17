// Crawler routes

package crawler

import (
	"github.com/go-chi/chi"
)

// Routes return the URLs registered in this module
func Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/fetch", FetchAllDataHandler)

	return router
}
