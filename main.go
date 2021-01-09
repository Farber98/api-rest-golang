package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	/* Cargamos la configuracion de rutas en la var router. */
	router := ConfiguroRouter()
	fmt.Println("Abrir servidor http://locahost:8080")
	/* Levantamos sv con config de rutas en el :8080 */
	servidor := http.ListenAndServe(":8080", router)
	log.Fatal(servidor)
}
