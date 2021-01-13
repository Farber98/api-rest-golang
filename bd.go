package main

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

/*ListarBD recibe bd, coleccion. Retorna error o resultados sort mas reciente. */
func ListarBD(db string, coleccion string) ([]Pelicula, error) {
	var resultados []Pelicula
	datosBD := Conexion().DB(db).C(coleccion)

	/* Resultados se ordenan mas recientes mediante el id */
	return resultados, datosBD.Find(nil).Sort("-_id").All(&resultados)
}

/*DameBD recibe bd, coleccion y IDpelicula. Retorna error o un resultado.  */
func DameBD(db string, coleccion string, OID bson.ObjectId) (Pelicula, error) {
	var resultados Pelicula
	return resultados, Conexion().DB(db).C(coleccion).FindId(OID).One(&resultados)
}

/*ModificarBD recibe bd, coleccion y IDpelicula. Retorna error o un resultado.  */
func ModificarBD(db string, coleccion string, OID bson.ObjectId, datosPelicula Pelicula) error {
	/* Encontramos el documento a modificar. */
	documento := bson.M{"_id": OID}
	/* $set indica el update en mongodb */
	cambio := bson.M{"$set": datosPelicula}
	return Conexion().DB(db).C(coleccion).Update(documento, cambio)
}

/*EliminarBD ....  */
func EliminarBD(db string, coleccion string, OID bson.ObjectId) error {
	/* Borramos el elemento del OID. */
	return Conexion().DB(db).C(coleccion).RemoveId(OID)
}
