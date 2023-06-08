package auth

import (
	"fmt"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hellor world", r.Header.Get("User-Agent"))
}
