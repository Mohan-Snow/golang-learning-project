package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"golang-learning-project/http-server/internal/model"
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
		return nil, &model.UserDataError{
			Message:   "Problem occurred while retrieving user by id.",
			Data:      userId,
			ErrorCode: 1004,
		}
	}
	return user, nil
}

func (s *Service) GetAll() ([]model.User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return nil, &model.GeneralError{
			Message:   "Problem occurred while retrieving all users.",
			ErrorCode: 1004,
		}
	}
	return users, nil
}

func (s *Service) SaveUser(user *model.User) error {
	err := s.repository.Save(user)
	if err != nil {
		return &model.UserDataError{
			Message:   "Problem occurred while saving new user.",
			Data:      fmt.Sprint(user),
			ErrorCode: 1003,
		}
	}
	return nil
}

func (s *Service) GenerateNames() ([]model.User, error) {
	response, err := http.Get(fmt.Sprintf(namesGenerationService, s.namesQueryParam))
	if err != nil {
		log.Println(err)
		return nil, &model.ExternalServiceError{
			Message:    "Problem occurred while retrieving names from Names Generation Service.",
			ServiceUrl: namesGenerationService,
			ErrorCode:  1001,
		}
	}
	var generatedNames []string
	err = json.NewDecoder(response.Body).Decode(&generatedNames)
	if err != nil {
		log.Println(err)
		return nil, &model.GeneralError{
			Message:   "Problem occurred while decoding response body.",
			ErrorCode: 1002,
		}
	}
	// save generated names to repository
	users, err := s.repository.SaveAll(generatedNames)
	if err != nil {
		log.Println(err)
		return nil, &model.UserDataError{
			Message:   "Problem occurred while saving multiple user names.",
			Data:      strings.Join(generatedNames, " "),
			ErrorCode: 1003,
		}
	}
	return users, nil
}
