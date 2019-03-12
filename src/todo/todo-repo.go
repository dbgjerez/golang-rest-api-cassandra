package todo

import (
	"encoding/json"
	"net/http"
)

func GetTodo(writer http.ResponseWriter, request *http.Request) {
	todo := Todo{1, "David"}
	json.NewEncoder(writer).Encode(todo)
}

func PostTodo(writer http.ResponseWriter, request *http.Request) {
	var t Todo
	json.NewDecoder(request.Body).Decode(&t)
}
