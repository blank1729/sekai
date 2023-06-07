package handlers

import (
	"fmt"
	"net/http"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hellor world", r.Header.Get("User-Agent"))
}
