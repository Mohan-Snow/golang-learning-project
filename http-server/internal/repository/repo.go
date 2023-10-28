package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"golang-learning-project/http-server/internal/model"
	"time"
)

type Repository struct {
	*sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db,
	}
}

// TODO: Review whole repository

func (r *Repository) FindAll() ([]model.User, error) {
	return r.getUserList()
}

func (r *Repository) FindUserById(id int) (*model.User, error) {
	var user model.User
	err := r.QueryRow("SELECT user_id, user_name FROM test_users WHERE user_id=$1", id).Scan(&user.Id, &user.Name)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *Repository) SaveAll(names []string) ([]model.User, error) {
	for i := 0; i < len(names); i++ {
		err := r.saveUser(names[i])
		if err != nil {
			// TODO: Find out how to handle multiple errors
		}
	}
	return r.getUserList()
}

func (r *Repository) Save(user *model.User) error {
	return r.saveUser(user.Name)
}

func (r *Repository) saveUser(username string) error {
	_, err := r.Exec("INSERT INTO test_users (user_name, created_on, changed_on) VALUES ($1,$2,$3)",
		username, time.Now(), time.Now())
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) getUserList() ([]model.User, error) {
	rows, err := r.Query("SELECT user_id, user_name FROM test_users")
	var users []model.User
	defer rows.Close()
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}
