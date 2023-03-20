package deck

import (
    "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"
)

func routes(router httprouter.Router) httprouter.Router {
    router.HandlerFunc("GET", "/", indexGet)
    return router
}

func indexGet(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Hola")
}