package main

import (
	"log"

	"gopkg.in/mgo.v2"
)

/*Conexion establece una conexion a la BD.  */
func Conexion() *mgo.Session {
	sesion, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		log.Fatal(err)
	}
	return sesion
}

/*InsertarBD recibe bd, coleccion y datos a insertar. */
func InsertarBD(db string, coleccion string, datos Pelicula) error {
	datosBD := Conexion().DB(db).C(coleccion)
	return datosBD.Insert(datos)
}

/*ListarBD recibe bd, coleccion */
/* func ListarBD(db string, coleccion string) *mgo.Query {
	datosBD := Conexion().DB(db).C(coleccion)
	return datosBD.Find(nil)
}*/
