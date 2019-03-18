package todo

import (
	"github.com/gocql/gocql"
	"log"
)

type LogLevel string

const (
	LOG_ERROR = "Error al guardar por: "
)
const (
	SELECT       = "SELECT id, text FROM todo"
	SELECT_BY_ID = "SELECT id, text FROM todo where id = ?"
	INSERT       = "INSERT INTO todo (id, text) VALUES (?, ?)"
	DELETE       = "DELETE from todo where id = ?"
)

func GetById(uuid gocql.UUID, session *gocql.Session) Todo {
	return getOne(uuid, session)
}

func GetTodo(session *gocql.Session) []Todo {
	return findAll(session)
}

func DeleteOne(id gocql.UUID, session *gocql.Session) {
	deleteOne(session, id)
}

func PostTodo(t *Todo, session *gocql.Session) {
	save(session, t)
}

func deleteOne(session *gocql.Session, id gocql.UUID) {
	session.Query(DELETE, id).Exec()
}

func getOne(id gocql.UUID, session *gocql.Session) Todo {
	var t Todo
	session.Query(SELECT_BY_ID, id).Scan(&t.ID, &t.Name);
	return t
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
	if err := session.Query(INSERT, id, todo.Name).Exec(); err != nil {
		log.Println(LOG_ERROR, err)
	}
}
