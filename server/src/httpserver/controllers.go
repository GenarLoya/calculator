package http_server

import (
	"encoding/json"
	"fmt"
	"net/http"

	pGen "genarold/calculator/src/postfixGenerator"
)

// Estructura para el cuerpo de la solicitud.
type bodyPostFix struct {
	Expression string `json:"exp"`
}

type responsePostFix struct {
	Result float64 `json:"result"`
}

func GetPostFixResult(w http.ResponseWriter, r *http.Request) {

	var body bodyPostFix
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&body); err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
		return
	}

	fmt.Println(body.Expression)

	// Generar el resultado de la expresión.
	result, err := pGen.PostfixManager(body.Expression)

	if err != nil {

		errMsg := fmt.Sprintf("Error al generar el resultado: %s", err.Error())

		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	// Convertir el resultado a JSON.
	jsonResult, err := json.Marshal(responsePostFix{Result: result})

	if err != nil {
		http.Error(w, "Error al convertir a JSON", http.StatusBadRequest)
		return
	}

	// Escribir el resultado en el cuerpo de la respuesta.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	w.Write(jsonResult)
}
