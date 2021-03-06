package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

/*Index es la funcion que se ejecuta al acceder a "/" */
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "El index pa.")
}

/*ListarPeliculas es la funcion que se ejecuta al acceder a "/peliculas" */
func ListarPeliculas(w http.ResponseWriter, r *http.Request) {

	/* Se obtiene resultados o error del metodo ListarBD.*/
	resultados, err := ListarBD("videoclub", "peliculas")

	/* Si es que existe, agarramos el error */
	if err != nil {
		log.Fatal(err)
	}

	Response(w, 200, resultados)
}

/*DamePelicula es la funcion que se ejecuta al acceder a "/peliculas/{id}" */
func DamePelicula(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) //Obtengo el id de la req en la url.
	IDPelicula := params["id"]

	/* Si ID no es hex, tira error y corta ejec. */
	if !bson.IsObjectIdHex(IDPelicula) {
		w.WriteHeader(404)
		return
	}

	/* Convertimos ID a hexa. */
	ObjectID := bson.ObjectIdHex(IDPelicula)

	resultados, err := DameBD("videoclub", "peliculas", ObjectID)

	/* Si hay error, cortamos ejec. */
	if err != nil {
		w.WriteHeader(404)
		return
	}

	Response(w, 200, resultados)
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

	/* Insertamos a la BD. */
	errorInsert := InsertarBD("videoclub", "peliculas", datosPelicula)
	if errorInsert != nil {
		/* En caso de que no sea exitoso, codigo error. */
		w.WriteHeader(500)
	}

	Response(w, 200, datosPelicula)
}

/*ModificarPelicula es la funcion que se ejecuta al acceder a "/peliculas/{id}" por PUT*/
func ModificarPelicula(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) //Obtengo el id de la req en la url.
	IDPelicula := params["id"]
	var datosPelicula Pelicula

	/* Si ID no es hex, tira error y corta ejec. */
	if !bson.IsObjectIdHex(IDPelicula) {
		w.WriteHeader(404)
		return
	}

	/* Convertimos ID a hexa. */
	ObjectID := bson.ObjectIdHex(IDPelicula)

	/* Decodificamos el json que nos llega en la r y lo bindeamos a datosPelicula. */
	decodificador := json.NewDecoder(r.Body)
	err := decodificador.Decode(&datosPelicula)

	/* En caso de que no pueda decodificar obtenemos error. */
	if err != nil {
		w.WriteHeader(500)
		return
	}

	/* Limpiamos el Body. */
	defer r.Body.Close()

	err = ModificarBD("videoclub", "peliculas", ObjectID, datosPelicula)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	Response(w, 200, datosPelicula)
}

/*EliminarPelicula es la funcion que se ejecuta al acceder a "/peliculas/{id}" por DELETE */
func EliminarPelicula(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) //Obtengo el id de la req en la url.
	IDPelicula := params["id"]
	var peliculaBorrada Pelicula

	/* Si ID no es hex, tira error y corta ejec. */
	if !bson.IsObjectIdHex(IDPelicula) {
		w.WriteHeader(404)
		return
	}

	/* Convertimos ID a hexa. */
	ObjectID := bson.ObjectIdHex(IDPelicula)

	/* Obtengo pelicula antes de borrar para mostrar nombre. */
	peliculaBorrada, errDame := DameBD("videoclub", "peliculas", ObjectID)
	if errDame != nil {
		w.WriteHeader(404)
		return
	}

	errElim := EliminarBD("videoclub", "peliculas", ObjectID)

	/* En caso de que no pueda borrar obtenemos error. */
	if errElim != nil {
		w.WriteHeader(404)
		return
	}

	/* Devolvemos mensaje como response. */
	resultado := Mensaje{"Exito", "La pelicula " + peliculaBorrada.Nombre + " ha sido borrada"}
	Response(w, 200, resultado)
}
