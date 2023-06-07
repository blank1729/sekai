package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HtmlParams(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	fmt.Fprintln(w, "you have requested", id)
}
