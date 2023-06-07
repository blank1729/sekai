package handler

import (
	"net/http"

	"github.com/blank1729/sekai/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// Basic CORS
// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
var corsmiddleware = cors.Handler(cors.Options{
	// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
	AllowedOrigins: []string{"https://*", "http://*"},
	// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	ExposedHeaders:   []string{"Link"},
	AllowCredentials: false,
	MaxAge:           300, // Maximum value not ignored by any of major browsers
})

func MainRouter(w http.ResponseWriter, r *http.Request) {
	// main router
	// **TODO** using the CORS for all the routes now, have to change it later
	router := chi.NewRouter()
	// adding middleware
	router.Use(middleware.Logger)
	router.Use(corsmiddleware)

	// using /api route
	apiRouter := chi.NewRouter()
	router.Mount("/api", apiRouter)

	// api routes
	router.Get("/api/", hello)

	// Products routes
	apiRouter.Get("/products", handlers.ProductsHandler)
	apiRouter.Get("/products/{id}", handlers.HtmlParams)

	// serving the request
	router.ServeHTTP(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
