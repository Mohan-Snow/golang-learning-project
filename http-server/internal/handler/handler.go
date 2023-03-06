package handler

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	"golang-learning-project/http-server/internal/model"
)

type Handler struct {
	service Service
}

type Service interface {
	GetUserById(id string) (*model.User, *model.Error)
	GetAll() map[int]string
	SaveUser(user *model.User) *model.User
	GenerateNames() (map[int]string, error)
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
		writeResponse(writer, http.StatusNotFound, err)
		return
	}
	writeResponse(writer, http.StatusOK, user)
}

func (h *Handler) Get(writer http.ResponseWriter, request *http.Request) {
	log.Println("Get all users")
	writeResponse(writer, http.StatusOK, h.service.GetAll())
}

func (h *Handler) Post(writer http.ResponseWriter, request *http.Request) {
	log.Println("Save user")
	user := model.User{}
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Print(err)
		writeResponse(writer, http.StatusBadRequest, model.Error{Error: "User was not saved. Internal service error."})
		return
	}
	h.service.SaveUser(&user)
	writeResponse(writer, http.StatusOK, user)
}

func (h *Handler) GenerateNames(writer http.ResponseWriter, request *http.Request) {
	log.Println("Generate names")
	names, err := h.service.GenerateNames()
	if err != nil {
		log.Print(err)
		writeResponse(writer, http.StatusInternalServerError, model.Error{Error: "Error during names generation. Internal service error."})
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
