package http_server

import (
	"fmt"
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("calculateROUTE")

		switch r.Method {
		case http.MethodPost:
			GetPostFixResult(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
