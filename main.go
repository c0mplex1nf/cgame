package main

import (
    "fmt"
    "net/http"
    "log"
    "github.com/julienschmidt/httprouter"
	"github.com/c0mplex1nf/cgame/pkg/deck"
)

func main() {

    log.Print("The Server have been started")
    router := httprouter.New()

    router = routes.routes.routes(router)

    err := http.ListenAndServe(":3000", router)
    log.Fatal(err)
}