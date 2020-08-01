// Hasta este punto el servidor es capaz de manejar diferentes rutas y asociarlas a diferentes handlers
// para esto, utilizamos mapas, etc, y al final lo logramos. Sin embargo está el caso si para
// especificos endpoints queremos validar si el usuario está logeado o no, etc, podriamos
// revisar en cada Handler si está logeado o no, pero por ejemplo si tuvieramos 3 handlers
// diferentes donde esto se hace tendríamos que poner la misma lógica, el mismo código en los 3,
// pero los Middlewares nos servirán para interceptar los requests y hacer toda esta lógica

package main

import (
	"encoding/json"
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MetaData interface{}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u)
}
