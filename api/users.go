package handler

import (
	"fmt"
	"net/http"
)

func users(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hellor world", r.Header.Get("User-Agent"))
}
