package middleware

import "github.com/go-chi/cors"

// Basic CORS
// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
var CORSMiddleware = cors.Handler(cors.Options{
	// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
	AllowedOrigins: []string{"https://*", "http://*"},
	// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
	AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	ExposedHeaders:   []string{"Link"},
	AllowCredentials: false,
	MaxAge:           300, // Maximum value not ignored by any of major browsers
})
