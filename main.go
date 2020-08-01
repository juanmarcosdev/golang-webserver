package main

// Para el server tendremos una mejor segmentación en nuestros archivos. Crearemos un nuevo archivo
// por cada cosa que vayamos a ocupando. Crearemos primero un archivo para definir la estructura
// general del servidor (server.go).

func main() {
	server := NewServer(":3000")
	server.Handle("GET", "/", HandleRoot)
	server.Handle("POST", "/create", PostRequest)
	server.Handle("POST", "/user", UserPostRequest)
	server.Handle("POST", "/api", server.AddMiddleware(HandleHome, CheckAuth(), Logging()))
	server.Listen()
}

// Anteriormente habíamos trabajado en el Struct de Server y prácticamente es capaz de crear
// una conexión y estar escuchando conexiones, sin embargo no se podían hacer ningún request todavía,
// debido a que no es capaz de "rutearlo" (routearlo), no es capaz de saber qué URL va a qué HANDLER
// Para ello creamos un nuevo archivo llamado router.go
