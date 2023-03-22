package main

import (
    "net/http"
    "log"
    "github.com/julienschmidt/httprouter"
	deck "github.com/c0mplex1nf/cgame/pkg/deck/infraestructure"
)

func main() {
    log.Print("The Server have been started")
    router := httprouter.New()
	router = deck.Routes(router)
    err := http.ListenAndServe(":3000", router)
    log.Fatal(err)
}