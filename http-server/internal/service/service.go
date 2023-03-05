package service

import (
	"golang-learning-project/http-server/internal/model"
	"strconv"
)

type Service struct {
	//data sync.Map
	repository map[int]string
}

func New() *Service {
	userRepo := make(map[int]string)
	return &Service{
		repository: userRepo,
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
