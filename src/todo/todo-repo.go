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
	TABLE        = "todo"
	FIELD_ID     = "id"
	FIELD_TEXT   = "text"
	SELECT       = "SELECT " + FIELD_ID + ", " + FIELD_TEXT + " FROM " + TABLE
	SELECT_BY_ID = "SELECT " + FIELD_ID + ", " + FIELD_TEXT + " FROM " + TABLE + " WHERE " + FIELD_ID + " = ?"
	INSERT       = "INSERT INTO " + TABLE + " (" + FIELD_ID + ", " + FIELD_TEXT + ") VALUES (?, ?)"
	DELETE       = "DELETE from " + TABLE + " WHERE " + FIELD_ID + " = ?"
	UPDATE       = "UPDATE " + TABLE + " SET " + FIELD_TEXT + " = ? WHERE " + FIELD_ID + " = ? IF EXISTS"
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

func UpdateOne(uuid gocql.UUID, todo *Todo, session *gocql.Session) {
	update(uuid, todo, session)
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

func update(uuid gocql.UUID, todo *Todo, session *gocql.Session) {
	session.Query(UPDATE, todo.Name, uuid).Exec()
}
