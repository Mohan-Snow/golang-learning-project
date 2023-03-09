package repository

import (
	_ "github.com/lib/pq"
	"golang-learning-project/http-server/internal/model"
)

type Repository struct {
	database map[int]string
}

func NewRepository(db map[int]string) *Repository {
	return &Repository{
		database: db,
	}
}

func (r *Repository) FindAll() []model.User {
	return getUserList(r.database)
}

func (r *Repository) FindUserById(id int) *model.User {
	userName, ok := r.database[id]
	if !ok {
		return nil
	}
	return &model.User{Id: id, Name: userName}
}

func (r *Repository) SaveAll(arr []string) []model.User {
	currLength := len(r.database)
	for i := 0; i < len(arr); i++ {
		currLength++
		r.database[currLength] = arr[i]
	}
	return getUserList(r.database)
}

func (r *Repository) Save(user *model.User) *model.User {
	r.database[user.Id] = user.Name
	return user
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
