package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

/* Creamos una variable de Peliculas, un slice que contiene objetos Pelicula. */
var peliculas = Peliculas{
	Pelicula{"Soul", 2020, "Yona Tan"},
	Pelicula{"Habia una vez", 2015, "Un Bru"},
	Pelicula{"Hey Ya", 1995, "Yayaya"}}

/*Index es la funcion que se ejecuta al acceder a "/" */
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "El index pa.")
}

/*ListarPeliculas es la funcion que se ejecuta al acceder a "/peliculas" */
func ListarPeliculas(w http.ResponseWriter, r *http.Request) {
	/* Indicamos que escriba w y que los datos vienen de slice peliculas.*/
	json.NewEncoder(w).Encode(peliculas)
}

/*DamePelicula es la funcion que se ejecuta al acceder a "/peliculas/{id}" */
func DamePelicula(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) //Obtengo el id de la req en la url.
	IDPelicula := params["id"]
	fmt.Fprintf(w, "Pelicula numero %s seleccionada", IDPelicula)
}

/*AgregarPelicula es la funcion que se ejecuta al acceder a "/pelicula" */
func AgregarPelicula(w http.ResponseWriter, r *http.Request) {
	/* Recibe los parametros que me llega por POST en la req */
	decoder := json.NewDecoder(r.Body)
	/* Convertimos lo recibido en algo que podamos utilizar */
	var datosPelicula Pelicula
	/* Agarramos el error en caso de haber uno*/
	err := decoder.Decode(&datosPelicula) // Como todavia no tiene nada ponemos &
	if err != nil {
		panic(err)
	}
	/* Limpiamos y cerramos la lectura con defer y close.*/
	defer r.Body.Close()
	log.Println(datosPelicula)
	/*Sobreescribimos el slice agregando la nueva pelicula  */
	peliculas = append(peliculas, datosPelicula)
}
