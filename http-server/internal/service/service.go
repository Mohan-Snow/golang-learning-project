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
	repository      map[int]string
	namesQueryParam string
}

func NewService(param string) *Service {
	userRepo := make(map[int]string)
	return &Service{
		repository:      userRepo,
		namesQueryParam: param,
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

func (s *Service) GetAll() []model.User {
	// TODO: Is it ok?
	return getUserList(s.repository)
}

func (s *Service) SaveUser(user *model.User) *model.User {
	s.repository[user.Id] = user.Name
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
	currLength := len(s.repository)
	for i := 0; i < 10; i++ {
		currLength++
		s.repository[currLength] = arr[i]
	}
	// TODO: Is it ok?
	return getUserList(s.repository), nil
}

// working with a copy of map here
// TODO: Is it ok?
func getUserList(users map[int]string) []model.User {
	var nameList []model.User
	for key, val := range users {
		user := model.User{Id: key, Name: val}
		nameList = append(nameList, user)
	}
	return nameList
}
