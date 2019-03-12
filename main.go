package main

import (
	"github.com/dbgjerez/go-todo-rest-api-cassandra/src/todo"
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
	PathPost   = "/todo"
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
	router.HandleFunc(PathGetAll, todo.GetTodo).Methods(GET)
	router.HandleFunc(PathPost, todo.PostTodo).Methods(POST)

	log.Fatal(http.ListenAndServe(":8000", router))
}
