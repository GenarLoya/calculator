package http_server

import (
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/calculate", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			GetPostFixResult(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
