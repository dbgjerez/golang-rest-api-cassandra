package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetTodo(writer http.ResponseWriter, request *http.Request) {
	todo := Todo{1, "David"}
	log.Println(DEBUG, todo)
	json.NewEncoder(writer).Encode(todo)
}

func PostTodo(writer http.ResponseWriter, request *http.Request) {
	var t Todo
	json.NewDecoder(request.Body).Decode(&t)
	log.Println(DEBUG, t)
}
