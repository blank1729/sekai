package handler

import (
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("something to search"))
	}
}
