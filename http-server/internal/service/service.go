package service

import (
	"encoding/json"
	"fmt"
	"golang-learning-project/http-server/internal/model"
	"net/http"
	"strconv"
)

const namesGenerationService = "https://names.drycodes.com/10?nameOptions=%s"

type Service struct {
	repository      Repository
	namesQueryParam string
}

type Repository interface {
	FindAll() []model.User
	FindUserById(id int) *model.User
	Save(user *model.User) *model.User
	SaveAll(arr []string) []model.User
}

func NewService(param string, r Repository) *Service {
	return &Service{
		repository:      r,
		namesQueryParam: param,
	}
}

func (s *Service) GetUserById(userId string) (*model.User, *model.Error) {
	id, _ := strconv.ParseInt(userId, 10, 64)
	user := s.repository.FindUserById(int(id))
	if &user == nil {
		return nil, &model.Error{Error: "User was not found!"}
	}
	return user, nil
}

func (s *Service) GetAll() []model.User {
	return s.repository.FindAll()
}

func (s *Service) SaveUser(user *model.User) *model.User {
	user = s.repository.Save(user)
	return user
}

func (s *Service) GenerateNames() ([]model.User, error) {
	response, err := http.Get(fmt.Sprintf(namesGenerationService, s.namesQueryParam))
	if err != nil {
		return nil, err
	}
	var arr []string
	err = json.NewDecoder(response.Body).Decode(&arr)
	if err != nil {
		return nil, err
	}
	// save generated names to repository
	users := s.repository.SaveAll(arr)
	return users, nil
}
