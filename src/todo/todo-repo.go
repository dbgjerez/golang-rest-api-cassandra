package todo

import (
	"encoding/json"
	"github.com/gocql/gocql"
	"log"
	"net/http"
)

const (
	LOG_ERROR = "Error al guardar por: "
)
const (
	SELECT = "SELECT id, text FROM todo"
	INSERT = "INSERT INTO todo (id, text) VALUES (?, ?)"
)

func GetTodo(writer http.ResponseWriter, request *http.Request, session *gocql.Session) {
	todo := findAll(session)
	json.NewEncoder(writer).Encode(&todo)
}

func PostTodo(writer http.ResponseWriter, request *http.Request, session *gocql.Session) {
	var t Todo
	json.NewDecoder(request.Body).Decode(&t)
}

func findAll(session *gocql.Session) Todo {
	var t Todo
	if err := session.Query(SELECT).PageSize(2).Scan(&t.ID, &t.Name); err != nil {
		log.Println(LOG_ERROR, err)
	}
	return t
}

func save(session *gocql.Session, todo *Todo) {
	if err := session.Query(INSERT,
		todo.ID, todo.Name).Exec(); err != nil {
		log.Println(LOG_ERROR, err)
	}
}
