package service

import (
	"golang-learning-project/http-server/internal/model"
	"strconv"
)

type Service struct {
	//data sync.Map
	repository map[int]string
}

func New(r map[int]string) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) GetNames() map[int]string {
	return map[int]string{0: "Mohan", 1: "Dinos", 2: "Rex"}
}

func (s *Service) GetUserById(userId string) (*model.User, *model.Error) {
	// username := fmt.Sprintf("%s", userName)
	id, _ := strconv.ParseInt(userId, 10, 64)
	userName := s.repository[int(id)]
	if userName == "" {
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
