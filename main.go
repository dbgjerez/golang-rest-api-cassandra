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

	s := todo.InitCluster()

	router := mux.NewRouter()
	router.HandleFunc(PathGetAll, func(writer http.ResponseWriter, request *http.Request) {
		todo.GetTodo(writer, request, s)
	}).Methods(GET)
	router.HandleFunc(PathPost, func(writer http.ResponseWriter, request *http.Request) {
		todo.PostTodo(writer, request, s)
	}).Methods(POST)

	log.Fatal(http.ListenAndServe(":8000", router))
}
