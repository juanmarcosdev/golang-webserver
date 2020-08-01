package main

import "net/http"

// Al decirle esto, le suscribimos al paquete de main. Y esto significa que el archivo main (principal)
// será capaz de leer todo lo que está en este archivo, e igual si creamos varios archivos en el mismo
// package, todos se van a conocer y van a poder utilizar lo que se van definiendo entre ellos.

// Creamos struct Server, y para que funcione debe tener un puerto en el cual está escuchando las
// conexiones

type Server struct {
	port   string
	router *Router // Creamos referencia a un Router
}

// Agregaremos las funciones que nos permitan ir creando el Server. Lo primero es la función
// NewServer, que nos permite instanciar un servidor como tal y que sea capaz de empezar a escuchar
// las conexiones. Al hacer esto estamos haciendo un componente altamente reutilizable en otros
// paquetes u otros proyectos.

// NewServer recibe el puerto en el cual se está escuchando, y debe devolver (recordando la clase
// de apuntadores) es como tal el Servidor (*Server, Server siendo una dirección de memoria y el
// apunta a ella y así no tenemos una copia sino un Server fidedigno), por ello utilizamos asterisco
// porque al hacer el return debemos utilizar & para que se refiera al valor que de verdad se desea
// modificar y no a copias que se estén generando.

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

// Con esta función somos capaces de crear un Servidor y poner manipularlo.

func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	_, exist := s.router.rules[path]
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

// Función Middleware

func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

// Necesitamos que nuestro Server esté escuchando conexiones, ahora haremos una receiver function
// de Server:

func (s *Server) Listen() error {
	// La barra "/" es la raíz inicial, el punto de entrada a la app
	http.Handle("/", s.router)
	err := http.ListenAndServe(s.port, nil) // Segundo parámetro el handler (estaremos pasando los propios entonces aquí pasamos nil)
	if err != nil {
		return err
	}
	return nil // Si todo sale bien retorna nil
}

// Ahora modificaremos Listen para que el Router sea quien tome las URL's y las procese como se debe.
