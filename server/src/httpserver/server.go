package http_server

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"

	"github.com/rs/cors"
)

func New(addr string) *http.Server {
	initRoutes()

	fmt.Println(color.GreenString("Server running on: "), color.CyanString(addr+"\n"))

	// Configura el manejador de CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Cambia esto a los dominios permitidos
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "Authorization", "Content-Type"},
		Debug:          true,
	})

	// Crea un nuevo manejador HTTP con CORS habilitado
	handler := c.Handler(http.DefaultServeMux)

	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
