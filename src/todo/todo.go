package todo

import "github.com/gocql/gocql"

type Todo struct {
	ID   gocql.UUID `json:"id,omitempty"`
	Name string     `json:"name,omitempty"`
}
