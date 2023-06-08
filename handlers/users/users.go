package users

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hellor world from /user", r.Header.Get("User-Agent"))
}
