package service

import (
	"time"

	"github.com/mbvisti/integration-test-in-go/entity"
)

type userStorage interface {
	CreateUser(newUser entity.User, createdAt time.Time) error
}

type User struct {
	repo userStorage
}

func NewUser(repo userStorage) *User {
	return &User{
		repo: repo,
	}
}

func (u User) New(name, sex, weightGoal, email string, age, height,
	activityLevel int) error {
	newUser, err := entity.NewUser(name, sex, weightGoal, email, age, height,
		activityLevel)
	if err != nil {
		return err
	}

	return u.repo.CreateUser(*newUser, time.Now())
}
