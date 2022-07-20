package controllers

import (
	"DrawFlowApp/services"

	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi"
)

type ProgramController struct {
	program_service services.ProgramService
}

func (rs ProgramController) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", rs.Index)
	r.Post("/save", rs.Store)
	r.Post("/execute", rs.Execute)

	r.Route("/{id}", func(r chi.Router) {
		r.Use(ProgramCtx)
		r.Get("/", rs.Show)
	})

	return r
}

func (rs ProgramController) Index(w http.ResponseWriter, r *http.Request) {

	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	response, err := rs.program_service.Index()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	programsJson, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(programsJson); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rs ProgramController) Store(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = rs.program_service.Store(body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write([]byte("Se ha accedido correctamente a la base de datos")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rs ProgramController) Execute(w http.ResponseWriter, r *http.Request) {

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	response, err := rs.program_service.Execute(body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	responseBytes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if _, err := w.Write(responseBytes); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rs ProgramController) Show(w http.ResponseWriter, r *http.Request) {
	programId := chi.URLParam(r, "id")

	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	response, err := rs.program_service.Show(programId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	programJson, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(programJson); err != nil {
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
