package main

import (
	"encoding/json"
	"net/http"
)

/*Response recibe el writer, status code y resultados. Retorna el JSON. */
func Response(w http.ResponseWriter, status int, resultados interface{}) {
	/* Escribimos la respuesta exitosa ya que no hay error. */
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200) // Codigo de exito
	/* Mostramos resultados en JSON. */
	json.NewEncoder(w).Encode(resultados)
}
