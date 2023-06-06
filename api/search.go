package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		wd, _ := os.Getwd()
		filePath := filepath.Join(wd, "data", "items.json")
		item, err := os.Open(filePath)
		if err != nil {
			fmt.Fprintln(w, err)
		}
		data, _ := io.ReadAll(item)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, data)
	}
}
