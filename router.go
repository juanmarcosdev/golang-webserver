package main

import (
	"net/http"
)

// Creamos tipo Router, que lleva adentro reglas. ¿Qué significa Reglas?, recuerda que nuestro Servidor
// tiene bien definido qué rutas es capaz de manejar, imaginate que viene una dirección que no existe
// o no está siendo manejada, es donde se ve el famoso error 404, entonces es muy importante dejar
// muy claro y bien definido de qué ruta pasamos a qué handler. Al final esto es un mapa,
// un mapa que pase de strings (rutas) a handlers (http)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

// Una vez teniendo esto listo haremos algo similar a lo que se hizo en el server: NewRouter

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// Qué se debe poner en las reglas? Recuerda que el servidor ya tenía definido un puerto, osea que
// teníamos un parámetro, para este caso no, Router empezará en un estado vacío, y tenemos que
// empezar en un mapa vacío. Al crear esto somos capaces de manejar nuestras rutas.

// Para pertenecer al club de Handlers en Server (poder ser tomado en cuenta como miembro de la interfaz)
// debemos crear función receiver que se llame ServeHTTP. El escritor w http.ResponseWriter nos permite
// responder a los requests

func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, methodExist := r.rules[path][method]
	return handler, methodExist, exist
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, methodExist, exist := r.FindHandler(request.URL.Path, request.Method)

	if !exist {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if !methodExist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)
}
