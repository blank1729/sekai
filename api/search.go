package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		item, err := os.Open("api/data/items.json")
		if err != nil {
			fmt.Fprintln(w, err)
		}
		data, _ := io.ReadAll(item)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, data)
	}
}
