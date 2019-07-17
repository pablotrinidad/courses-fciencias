// Courses F. Ciencias
// This project fetches all available course schedules from UNAM's
// faculty of science website.

package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"

	"github.com/pablotrinidad/courses-fciencias/api/crawler"
)

// Routes return an HTTP chi router
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		// Set Content-Type headers as application/json
		render.SetContentType(render.ContentTypeJSON),

		// Log API request calls
		middleware.Logger,

		// Redirect slashes to no slush URL versions
		middleware.RedirectSlashes,

		// Recover from panics without crashing server
		middleware.Recoverer,
	)

	router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/crawler", crawler.Routes())
	})
	return router
}

func main() {
	router := Routes()

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s\n", method, route) // Walk and print out all routes
		return nil
	}
	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Logging err %s\n", err.Error()) // Panic if there's an error
	}

	log.Fatal(http.ListenAndServe(":5000", router))
}
