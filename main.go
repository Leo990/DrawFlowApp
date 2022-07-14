package main

import (
	"DrawFlowApp/views"

	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {

	port := "8080"

	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	log.Printf("Starting up on http://localhost:%s", port)

	r := chi.NewRouter()

	r.Mount("/programs", views.ProgramResources{}.Routes())

	log.Fatal(http.ListenAndServe(":"+port, r))

}
