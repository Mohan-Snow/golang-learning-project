package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strconv"
)

type UserHandler struct {
	//data sync.Map
	data map[int]string
}

func getNames(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get Names")
	names := map[int]string{0: "Mohan", 1: "Dinos", 2: "Some Name"}
	writeResponse(writer, http.StatusOK, names)
	//for key, val := range names {
	//	temp := fmt.Sprintf("Id=%d Name=%s", key, val)
	//	fmt.Println(temp)
	//	TODO: Handle error
	//fmt.Fprintln(writer, temp)
	//}
}

func getUserById(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	log.Printf("Get User by id=%s", id)
	username := fmt.Sprintf("NewUser[%s]", id)
	userId, _ := strconv.ParseInt(id, 10, 64)
	user := User{int(userId), username}
	writeResponse(writer, http.StatusOK, user)
}

func (h *UserHandler) Get(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get all users")
	writeResponse(writer, http.StatusOK, h.data)
}

func (h *UserHandler) Post(writer http.ResponseWriter, request *http.Request) {
	log.Println("Save user")
	user := User{}
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		writeResponse(writer, http.StatusBadRequest, Error{err.Error()})
	}
	h.data[user.Id] = user.Name
	writeResponse(writer, http.StatusOK, user)
}

func writeResponse(writer http.ResponseWriter, code int, v interface{}) {
	body, _ := json.Marshal(v)
	writer.WriteHeader(code)
	writer.Write(body)
}
