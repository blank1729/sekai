package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Item struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fd, err := os.Open("data/items.json")
		if err != nil {
			fmt.Fprintln(w, err)
		}
		data, _ := io.ReadAll(fd)
		var items []Item
		json.Unmarshal([]byte(data), &items)

		w.Header().Set("Content-Type", "application/json")
		// fmt.Fprintln(w, data)
		json.NewEncoder(w).Encode(items)
	}
}
