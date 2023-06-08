package orders

import (
	"fmt"
	"net/http"
)

func OrdersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hellor world", r.Header.Get("User-Agent"))
}
