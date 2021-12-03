package service

import (
	"time"

	"github.com/mbvisti/integration-test-in-go/entity"
)

type weightStorage interface {
	CreateWeightEntry(newWeight entity.Weight, createdAt time.Time) error
}

type Weight struct {
	repo weightStorage
}

func NewWeight(repo weightStorage) *Weight {
	return &Weight{
		repo: repo,
	}
}

func (w Weight) New(weight, userID, bmr, dailyCaloricIntake int) error {
	newWeight, err := entity.NewWeight(weight, userID, bmr, dailyCaloricIntake)
	if err != nil {
		return err
	}

	return w.repo.CreateWeightEntry(*newWeight, time.Now())
}
