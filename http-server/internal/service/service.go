package service

import (
	"encoding/json"
	"golang-learning-project/http-server/internal/model"
	"net/http"
	"strconv"
)

const namesGenerationService = "https://names.drycodes.com/10"

type Service struct {
	//data sync.Map
	repository map[int]string
	token      string
}

func NewService(externalApiToken string) *Service {
	userRepo := make(map[int]string)
	return &Service{
		repository: userRepo,
		token:      externalApiToken,
	}
}

func (s *Service) GetUserById(userId string) (*model.User, *model.Error) {
	id, _ := strconv.ParseInt(userId, 10, 64)
	userName, ok := s.repository[int(id)]
	if !ok {
		return nil, &model.Error{Error: "User was not found!"}
	}
	return &model.User{Id: int(id), Name: userName}, nil
}

func (s *Service) GetAll() map[int]string {
	return s.repository
}

func (s *Service) SaveUser(user *model.User) *model.User {
	s.repository[user.Id] = user.Name
	return user
}

func (s *Service) GenerateNames() (map[int]string, error) {
	response, err := http.Get(namesGenerationService)
	if err != nil {
		return nil, err
	}
	var arr []string
	err = json.NewDecoder(response.Body).Decode(&arr)
	if err != nil {
		return nil, err
	}
	for i := 0; i < 10; i++ {
		s.repository[i+1] = arr[i]
	}
	return s.repository, nil
}
