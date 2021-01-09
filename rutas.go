package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

/*Ruta define el modelo de una ruta. */
type Ruta struct {
	Nombre     string
	Metodo     string
	Patron     string
	HandleFunc http.HandlerFunc
}

/*Rutas define slice de elementos ruta.*/
type Rutas []Ruta

/* Definimos nuestro array de rutas con cada ruta que posee la APP. */
var rutas = Rutas{
	Ruta{"Index", "GET", "/", Index},
	Ruta{"ListarPeliculas", "GET", "/peliculas", ListarPeliculas},
	Ruta{"DamePelicula", "GET", "/peliculas/{id}", DamePelicula}}

/*ConfiguroRouter devuelve la configuracion de rutas en var router.  */
func ConfiguroRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	/* Iteramos slice rutas y seteamos nuestro router */
	for _, ruta := range rutas {
		router.
			Name(ruta.Nombre).
			Methods(ruta.Metodo).
			Path(ruta.Patron).
			Handler(ruta.HandleFunc)
	}
	return router
}
