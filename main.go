package main

import (
	"encoding/json"
	"github.com/dbgjerez/go-todo-rest-api-cassandra/src/todo"
	"github.com/gocql/gocql"
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
	PathGetAll  = "/todo"
	PathPost    = "/todo"
	PathDelete  = "/todo/{id}"
	PathGetById = "/todo/{id}"
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
	router.HandleFunc(PathGetAll, getAll(s)).Methods(GET)
	router.HandleFunc(PathGetById, getById(s)).Methods(GET)
	router.HandleFunc(PathPost, post(s)).Methods(POST)
	router.HandleFunc(PathDelete, delete(s)).Methods(DELETE)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getById(s *gocql.Session) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var id gocql.UUID
		vars := mux.Vars(request)
		id, _ = gocql.ParseUUID(vars["id"])
		t := todo.GetById(id, s)
		json.NewEncoder(writer).Encode(&t)
	}
}

func delete(s *gocql.Session) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		var id gocql.UUID
		vars := mux.Vars(request)
		id, _ = gocql.ParseUUID(vars["id"])
		todo.DeleteOne(id, s)
	}
}

func getAll(s *gocql.Session) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		log.Println(DEBUG, LOG_GET_ALL)
		res := todo.GetTodo(s)
		json.NewEncoder(writer).Encode(&res)
	}
}

func post(s *gocql.Session) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		t := read(request)
		log.Println(DEBUG, LOG_POST, t)
		todo.PostTodo(&t, s)
		writer.WriteHeader(200)
	}
}

func read(r *http.Request) todo.Todo {
	var t todo.Todo
	json.NewDecoder(r.Body).Decode(&t)
	return t
}
