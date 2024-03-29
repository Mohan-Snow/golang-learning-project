package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	"golang-learning-project/http-server/internal/model"
)

type Handler struct {
	service Service
}

type Service interface {
	GetUserById(id string) (*model.User, error)
	GetAll() ([]model.User, error)
	SaveUser(user *model.User) error
	GenerateNames() ([]model.User, error)
}

func NewHandler(s Service) *Handler {
	return &Handler{
		service: s,
	}
}

func (h *Handler) GetUserById(writer http.ResponseWriter, request *http.Request) {
	userId := chi.URLParam(request, "id")
	log.Printf("Get User by id=%s", userId)
	user, err := h.service.GetUserById(userId)
	if err != nil {
		log.Println(err)
		writeResponse(writer, http.StatusNotFound, err)
		return
	}
	writeResponse(writer, http.StatusOK, user)
}

func (h *Handler) Get(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get all users")
	users, err := h.service.GetAll()
	if err != nil {
		log.Println(err)
		writeResponse(writer, http.StatusOK, err)
		return
	}
	if len(users) == 0 {
		writeResponse(writer, http.StatusOK, "No users were found")
		return
	}
	writeResponse(writer, http.StatusOK, users)
}

func (h *Handler) Post(writer http.ResponseWriter, request *http.Request) {
	log.Println("Save user")
	user := model.User{}
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Print(err)
		writeResponse(writer, http.StatusBadRequest,
			model.GeneralError{
				Message:   "User was not saved. Problem occurred while decoding request body data.",
				ErrorCode: 1002,
			})
		return
	}
	err = h.service.SaveUser(&user)
	if err != nil {
		log.Println(err)
		writeResponse(writer, http.StatusBadRequest, err)
		return
	}
	writeResponse(writer, http.StatusOK, fmt.Sprintf("User %s was succsessfully saved.", user.Name))
}

func (h *Handler) GenerateNames(writer http.ResponseWriter, request *http.Request) {
	log.Println("Generate names")
	names, err := h.service.GenerateNames()
	if err != nil {
		log.Print(err)
		writeResponse(writer, http.StatusInternalServerError, err)
		return
	}
	writeResponse(writer, http.StatusOK, names)
}

func writeResponse(writer http.ResponseWriter, code int, v interface{}) {
	body, _ := json.Marshal(v)
	writer.WriteHeader(code)
	_, err := writer.Write(body)
	if err != nil {
		log.Print(err) // logging with log package because fmt package is not concurrent safe
	}
}
