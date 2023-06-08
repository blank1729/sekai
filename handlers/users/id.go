package users

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func IdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Fprintf(w, "you have requested user id %s", id)
}
