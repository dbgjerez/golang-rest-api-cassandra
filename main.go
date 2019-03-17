package main

import (
	"encoding/json"
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
	LOG_POST    = "Creando todo: "
	LOG_GET_ALL = "Buscando todos los todo"
	LOG_GET_ONE = "Recuperando todo con id:"
	LOG_DELETE  = "Eliminando todo con id:"
)

const (
	PathGetAll = "/todo"
	PathPost   = "/todo"
	PathDelete = "/todo/{id}"
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
		log.Println(DEBUG, LOG_GET_ALL)
		todo.GetTodo(writer, request, s)
	}).Methods(GET)
	router.HandleFunc(PathPost, func(writer http.ResponseWriter, request *http.Request) {
		var todo *todo.Todo
		json.NewDecoder(request.Body).Decode(&todo)
		log.Println(DEBUG, LOG_POST, todo)
		todo.PostTodo(writer, todo, s)
	}).Methods(POST)
	router.HandleFunc(PathDelete, func(writer http.ResponseWriter, request *http.Request) {
		todo.DeleteOne(writer, request, s)
	}).Methods(DELETE)

	log.Fatal(http.ListenAndServe(":8000", router))
}
