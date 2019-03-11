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
	PathGetAll = "/todo"
	PathPost   = PathGetAll
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

	router.HandleFunc(PathGetAll, GetTodo).Methods(GET)
	router.HandleFunc(PathPost, PostTodo).Methods(POST)

	log.Fatal(http.ListenAndServe(":8000", router))
}
