package entity

import (
	"time"

	"github.com/pkg/errors"
	"gopkg.in/guregu/null.v4"
)

var (
	ErrNewWeight = errors.New("could not create new weight")
)

type Weight struct {
	ID                 int64
	CreatedAt          time.Time
	UpdatedAt          null.Time
	Weight             int
	UserID             int
	BMR                int
	DailyCaloricIntake int
}

func NewWeight(weight, userID, bmr, dailyCaloricIntake int) (*Weight, error) {
	newWeight := &Weight{
		Weight:             weight,
		UserID:             userID,
		BMR:                bmr,
		DailyCaloricIntake: dailyCaloricIntake,
	}

	if isInvalid := newWeight.IsValid(); !isInvalid {
		// returning errors with the stack trace gives us some super nice
		// debugging options
		return nil, errors.WithStack(ErrNewWeight)
	}
	return newWeight, nil
}

// IsValid: by creating this method we ensure that our business rules regarding
// creating a new user is always satified, as long as we use the NewUser
// method
func (w Weight) IsValid() bool {
	if w.Weight == 0 {
		return false
	}
	if w.UserID == 0 {
		return false
	}
	if w.BMR == 0 {

		return false
	}
	if w.DailyCaloricIntake == 0 {

		return false
	}
	return true
}
