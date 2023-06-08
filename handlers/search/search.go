package search

import (
	"net/http"
	"os/exec"
)

func RootWithQueryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// checking if files are copied
		cmd := exec.Command("ls", "-al")
		output, _ := cmd.Output()
		w.Write(output)
		w.Write([]byte("something to search"))
	}
}
