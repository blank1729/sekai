package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Item struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fd, err := os.Open("data/items.json")
		if err != nil {
			log.Fatal("unable to open the file")
		}
		data, _ := io.ReadAll(fd)
		var items []Item
		err = json.Unmarshal([]byte(data), &items)
		if err != nil {
			fmt.Println("unable to unmarshal", err)
		}
		jsonData, err := json.MarshalIndent(items, "  ", "  ")
		if err != nil {
			fmt.Println("unable to marshal", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
