package handlers

import (
	"fmt"
	"net/http"
)

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hellor world", r.Header.Get("User-Agent"))
}
