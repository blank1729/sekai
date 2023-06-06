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
			fmt.Fprintln(w, err)
		}
		data, _ := io.ReadAll(fd)
		var items []Item
		err = json.Unmarshal([]byte(data), &items)
		if err != nil {
			fmt.Println("unable to unmarshal", err)
		}
		w.Header().Set("Content-Type", "application/json")
		jsonData, err := json.Marshal(items)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(jsonData)
	}
}
