package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type LogLevel string

const (
	DEBUG LogLevel = "DEBUG"
	INFO  LogLevel = "INFO"
	ERROR LogLevel = "ERROR"
)

const (
	PATH = "/todo"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

func main() {
	log.Println(INFO, "Servidor iniciado")
	router := mux.NewRouter()

	router.HandleFunc(PATH, GetTodo).Methods(GET)

	log.Fatal(http.ListenAndServe(":8000", router))
}
