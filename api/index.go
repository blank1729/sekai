package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func MainRouter(w http.ResponseWriter, r *http.Request) {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/api/", hello)
	router.ServeHTTP(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
