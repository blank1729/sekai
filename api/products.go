package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var Product struct {
	id    string
	name  string
	price int
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		item, _ := os.Open("data/item.json")
		data, _ := io.ReadAll(item)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, data)
	}
}
