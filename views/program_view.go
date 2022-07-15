package views

import (
	"DrawFlowApp/services"
	"fmt"

	"context"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

type ProgramResources struct {
	program_service services.ProgramService
}

func (rs ProgramResources) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/save", rs.Store) // POST /posts - Create a new post.
	r.Post("/execute", rs.Execute)
	// r.Get("/", rs.Index) // GET /posts - Read a list of posts.

	// r.Route("/{id}", func(r chi.Router) {
	// 	r.Use(ProgramCtx)
	// 	r.Get("/", rs.Show)   // GET /posts/{id} - Read a single post by :id.
	// 	r.Put("/", rs.Update) // PUT /posts/{id} - Update a single post by :id.
	// })

	return r
}

func (rs ProgramResources) Store(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	errors := rs.program_service.Store(body)

	if errors != nil {
		http.Error(w, errors.Error(), http.StatusInternalServerError)
	}

	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rs ProgramResources) Execute(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response := rs.program_service.Execute(body)

	fmt.Println(response)

	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := io.Copy(w, r.Body); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func ProgramCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), "id", chi.URLParam(r, "id"))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// func (rs ProgramResources) Update(w http.ResponseWriter, r *http.Request) {

// 	resp, err := rs.program_service.Index()

// 	fmt.Println(resp, err)
// }

// func (rs ProgramResources) Show(w http.ResponseWriter, r *http.Request) {

// 	resp, err := rs.program_service.Index()

// 	fmt.Println(resp, err)
// }
// Request Handler - GET /posts - Read a list of posts.
// func (rs ProgramResources) Index(w http.ResponseWriter, r *http.Request) {
// 	//resp, err := rs.program_service.Index()

// 	body, err := io.ReadAll(r.Body)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 	}

// 	var program schemas.Program

// 	json.Unmarshal(body, &program)

// 	fmt.Println(program)

// 	defer r.Body.Close()

// 	w.Header().Set("Content-Type", "application/json")

// 	if _, err := io.Copy(w, r.Body); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// }
