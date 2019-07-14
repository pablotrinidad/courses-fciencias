// Course FCiencias.
// A web crawler made to download and store UNAM's Faculty of Science
// available courses schedules.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/pablotrinidad/courses-fciencias/fciencias"
)

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	majors := fciencias.GetMajors()
	data, _ := json.MarshalIndent(majors, "", "    ")
	fmt.Fprintf(w, "%s\n", data)
}
