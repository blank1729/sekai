package handler

import (
	"net/http"
	"os/exec"

	"github.com/blank1729/sekai/utils"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// checking if files are copied
		cmd := exec.Command("ls", "-al")
		output, _ := cmd.Output()
		w.Write(output)
		w.Write([]byte("something to search"))

		// utils func
		utils.Print()
	}
}
