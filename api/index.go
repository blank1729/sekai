package handler

import (
	"net/http"

	authHandlers "github.com/blank1729/sekai/handlers/auth"
	productsHandler "github.com/blank1729/sekai/handlers/products"
	searchHandlers "github.com/blank1729/sekai/handlers/search"
	usersHandlers "github.com/blank1729/sekai/handlers/users"
	md "github.com/blank1729/sekai/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func MainRouter(w http.ResponseWriter, r *http.Request) {
	// main router
	// **TODO** using the CORS for all the routes now, have to change it later
	router := chi.NewRouter()
	// adding middleware
	router.Use(middleware.Logger)
	router.Use(md.CORSMiddleware)

	// using /api route
	apiRouter := chi.NewRouter()
	router.Mount("/api", apiRouter)
	apiRouter.Use(md.AuthMiddleware)

	// api routes
	router.Get("/api/", hello)

	// auth routes
	apiRouter.Get("/auth/login", authHandlers.LoginHandler)

	// users routes
	apiRouter.Get("/user", usersHandlers.RootHandler)
	apiRouter.Get("/user/{id}", usersHandlers.IdHandler)

	// products routes
	apiRouter.Get("/products", productsHandler.RootHandler)
	apiRouter.Get("/products/{id}", productsHandler.IdHandler)

	// search route
	apiRouter.Get("/search", searchHandlers.RootWithQueryHandler)

	// serving the request
	router.ServeHTTP(w, r)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("you are accessing the / off the api and there is nothing here to see, take care"))
}
