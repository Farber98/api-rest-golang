package main

/*Pelicula creara las instancias de Peliculas.
Ademas indicamos que nombre llevara en el json.*/
type Pelicula struct {
	Nombre   string `json:"nombre"`
	Fecha    int    `json:"fecha"`
	Director string `json:"director"`
}

/*Peliculas sera un slice de objetos Pelicula.  */
type Peliculas []Pelicula
