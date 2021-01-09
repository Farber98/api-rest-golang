package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

/*Index es la funcion que se ejecuta al acceder a "/" */
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "El index pa.")
}

/*ListarPeliculas es la funcion que se ejecuta al acceder a "/peliculas" */
func ListarPeliculas(w http.ResponseWriter, r *http.Request) {
	/* Creamos una variable de Peliculas, un slice que contiene objetos Pelicula. */
	peliculas := Peliculas{
		Pelicula{"Soul", 2020, "Yona Tan"},
		Pelicula{"Habia una vez", 2015, "Un Bru"},
		Pelicula{"Hey Ya", 1995, "Yayaya"}}
	/* Indicamos que escriba w y que los datos vienen de peliculas.*/
	json.NewEncoder(w).Encode(peliculas)
}

/*DamePelicula es la funcion que se ejecuta al acceder a "/peliculas/{id}" */
func DamePelicula(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) //Obtengo el id de la req en la url.
	IDPelicula := params["id"]
	fmt.Fprintf(w, "Pelicula numero %s seleccionada", IDPelicula)
}