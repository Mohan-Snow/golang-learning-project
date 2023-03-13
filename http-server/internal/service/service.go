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
	FindAll() ([]model.User, error)
	FindUserById(id int) (*model.User, error)
	Save(user *model.User) error
	SaveAll(arr []string) ([]model.User, error)
}

func NewService(param string, r Repository) *Service {
	return &Service{
		repository:      r,
		namesQueryParam: param,
	}
}

func (s *Service) GetUserById(userId string) (*model.User, error) {
	id, _ := strconv.ParseInt(userId, 10, 64)
	user, err := s.repository.FindUserById(int(id))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *Service) GetAll() ([]model.User, error) {
	return s.repository.FindAll()
}

func (s *Service) SaveUser(user *model.User) error {
	err := s.repository.Save(user)
	if err != nil {
		return err
	}
	return nil
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
	users, err := s.repository.SaveAll(arr)
	if err != nil {
		return nil, err
	}
	return users, nil
}
