// En las clases anteriores fuimos capaces de imprimir nuestro Hola Mundo desde el Server,
// realmente fue un Milestone que hemos alcanzado. Sin embargo, no es la forma correcta de manejarlo.
// Lo mejor es tener handlers por a parte que se encarguen de manejar todo esto. Para ello está
// este archivo handlers.go, que da una mejor segmentación de esto. La única responsabilidad del Router
// es saber que handler está asociado a qué ruta.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Recordemos cuando creamos el Router y fuimos capaces de imprimir Hola Mundo. Implementamos
// una función llamada ServeHTTP, que tenía dos cosas: un escritor para responder las peticiones
// de los clientes, y uno que se llamab Request, dicho request tiene todo lo que necesitamos para
// saber todos los parámetros, qué rutas se están enviando, tratando de acceder. Necesitaremos
// ambos para crear el Handler

// Lo que intentaremos hacer es imprimir el Hola Mundo del Router pero utilizando este Handler

// Handler ruta principal
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World desde el Handler!")
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the API Endpoint")
}

func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metadata MetaData
	err := decoder.Decode(&metadata)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	fmt.Fprintf(w, "Payload %v\n", metadata)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// Ya creado el Handler, no está mapeado a nada. Vamos a modificar el Router para que asocie
// Handler con Ruta.
