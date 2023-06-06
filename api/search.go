package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// wd, _ := os.Getwd()
		// filePath := filepath.Join(wd, "data", "items.json")
		item, err := os.Open("data/items.json")
		if err != nil {
			fmt.Fprintln(w, "can't open file")
		}
		cmd := exec.Command("ls", "-l")
		output, _ := cmd.Output()
		fmt.Fprintln(w, err, string(output))
		data, _ := io.ReadAll(item)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, data)
	}
}
