package products

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/blank1729/sekai/utils"
	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Price      string   `json:"price"`
	Categories []string `json:"categories"`
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	utils.Cors(w, r)

	// Handle preflight requests
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method == http.MethodGet {
		if r.URL.Query().Has("id") {
			itemHandler(w, r, r.URL.Query().Get("id"))
		} else {
			productHandler(w, r)
		}
		return
	}

}

func itemHandler(w http.ResponseWriter, r *http.Request, id string) {

	if !valid(id) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("invalid id"))
		return
	}

	fd, err := os.Open("data/items.json")
	if err != nil {
		log.Fatal("unable to open the file")
	}
	defer fd.Close()
	data, _ := io.ReadAll(fd)
	var items []Product
	err = json.Unmarshal([]byte(data), &items)
	if err != nil {
		fmt.Println("unable to unmarshal", err)
	}
	var found bool
	var out Product

	for _, item := range items {
		if item.Id == id {
			found = true
			out = item
			break
		}
	}

	if found {
		jsonData, err := json.MarshalIndent(out, "  ", "  ")
		if err != nil {
			fmt.Println("unable to marshal", err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Product not found"))
	}

}

func productHandler(w http.ResponseWriter, r *http.Request) {
	fd, err := os.Open("data/items.json")
	if err != nil {
		log.Fatal("unable to open the file")
	}
	defer fd.Close()
	data, _ := io.ReadAll(fd)
	var items []Product
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

func valid(itemId string) bool {
	validPattern := regexp.MustCompile(`^[0-9]+$`)
	return validPattern.MatchString(itemId)
}
