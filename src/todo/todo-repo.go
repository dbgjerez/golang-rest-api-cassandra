package todo

import (
	"encoding/json"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type LogLevel string

const (
	DEBUG LogLevel = "DEBUG"
)

const (
	LOG_ERROR = "Error al guardar por: "
)
const (
	SELECT = "SELECT id, text FROM todo"
	INSERT = "INSERT INTO todo (id, text) VALUES (?, ?)"
	DELETE = "DELETE from todo where id = ?"
)

func GetTodo(writer http.ResponseWriter, request *http.Request, session *gocql.Session) {
	todo := findAll(session)
	json.NewEncoder(writer).Encode(&todo)
}

func DeleteOne(writer http.ResponseWriter, request *http.Request, session *gocql.Session) {
	var id gocql.UUID
	vars := mux.Vars(request)
	id, _ = gocql.ParseUUID(vars["id"])
	deleteOne(session, id)
	writer.WriteHeader(200)
}

func (t Todo) PostTodo(writer http.ResponseWriter, todo *Todo, session *gocql.Session) {
	writer.WriteHeader(200)
	save(session, todo)
}

func deleteOne(session *gocql.Session, id gocql.UUID) {
	session.Query(DELETE, id)
}

func findAll(session *gocql.Session) []Todo {
	var ts []Todo
	var t Todo
	it := session.Query(SELECT).Iter()
	for it.Scan(&t.ID, &t.Name) {
		ts = append(ts, t)
	}
	if err := it.Close(); err != nil {
		log.Println(LOG_ERROR, err)
	}
	return ts
}

func save(session *gocql.Session, todo *Todo) {
	var id gocql.UUID = gocql.TimeUUID()
	if err := session.Query(INSERT,
		id, todo.Name).Exec(); err != nil {
		log.Println(LOG_ERROR, err)
	}
}
