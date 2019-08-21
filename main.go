package main

import (
	"encoding/json"
	"github.com/dbgjerez/go-todo-rest-api-cassandra/src/todo"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type LogLevel string

const (
	DEBUG LogLevel = "DEBUG"
	INFO  LogLevel = "INFO"
	ERROR LogLevel = "ERROR"
)

const (
	LOG_POST    = "Creado todo: "
	LOG_GET_ALL = "Buscando todos los todo "
	LOG_GET_ONE = "Recuperando todo con id: "
	LOG_DELETE  = "Eliminando todo con id: "
)

const (
	PathGetAll  = "/todo"
	PathPost    = "/todo"
	PathDelete  = "/todo/{id}"
	PathGetById = "/todo/{id}"
	PathPut     = "/todo/{id}"
)

const (
	GET               = "GET"
	POST              = "POST"
	PUT               = "PUT"
	DELETE            = "DELETE"
	HeaderContenttype = "Content-Type"
	HeaderIdSession   = "IdSession"
	ContentTypeJson   = "application/json"
)

func main() {
	log.Println(INFO, "Servidor iniciado a las "+time.Now().String())

	s := todo.InitCluster()

	router := mux.NewRouter()
	router.Use(commonMiddleware)

	router.HandleFunc(PathGetAll, getAll(s)).Methods(GET)
	router.HandleFunc(PathGetById, getById(s)).Methods(GET)
	router.HandleFunc(PathPut, put(s)).Methods(PUT)
	router.HandleFunc(PathPost, post(s)).Methods(POST)
	router.HandleFunc(PathDelete, delete(s)).Methods(DELETE)

	log.Fatal(http.ListenAndServe(":8000", router))
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add(HeaderContenttype, ContentTypeJson)
		next.ServeHTTP(w, r)
	})
}

func put(session *gocql.Session) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := extractId(request)
		t := read(request)
		todo.UpdateOne(id, &t, session)
	}
}

func getById(s *gocql.Session) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := extractId(request)
		log.Println(DEBUG, LOG_GET_ONE+id.String())
		t := todo.GetById(id, s)
		json.NewEncoder(writer).Encode(&t)
	}
}

func delete(s *gocql.Session) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		id := extractId(request)
		todo.DeleteOne(id, s)
	}
}

func extractId(request *http.Request) gocql.UUID {
	var id gocql.UUID
	vars := mux.Vars(request)
	id, _ = gocql.ParseUUID(vars["id"])
	return id
}

func getAll(s *gocql.Session) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		IdSession := request.Header.Get(HeaderIdSession)
		log.Println(DEBUG, LOG_GET_ALL+IdSession)
		res, state := todo.GetTodo(s, IdSession)
		writer.Header().Add(HeaderIdSession, string(state))
		json.NewEncoder(writer).Encode(&res)
	}
}

func post(s *gocql.Session) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		t := read(request)
		todo.PostTodo(&t, s)
		log.Println(DEBUG, LOG_POST, t)
		writer.WriteHeader(201)
	}
}

func read(r *http.Request) todo.Todo {
	var t todo.Todo
	json.NewDecoder(r.Body).Decode(&t)
	return t
}
