package entity

import "github.com/pkg/errors"

var (
	ErrNewUser = errors.New("could not create new user")
)

type User struct {
	ID            int
	Name          string
	Age           int
	Height        int
	Sex           string
	ActivityLevel int
	WeightGoal    string
	Email         string
}

func NewUser(name, sex, weightGoal, email string, age, height,
	activityLevel int) (*User, error) {
	newUser := &User{
		Name:          name,
		Age:           age,
		Height:        height,
		Sex:           sex,
		ActivityLevel: activityLevel,
		WeightGoal:    weightGoal,
		Email:         email,
	}

	if isInvalid := newUser.IsValid(); !isInvalid {
		// returning errors with the stack trace gives us some super nice
		// debugging options
		return nil, errors.WithStack(ErrNewUser)
	}
	return newUser, nil
}

// IsValid: by creating this method we ensure that our business rules regarding
// creating a new user is always satified, as long as we use the NewUser
// method
func (u User) IsValid() bool {
	if u.Name == "" {
		return false
	}
	if u.Age == 0 {
		return false
	}
	if u.Height == 0 {
		return false
	}
	if u.Sex == "" {
		return false
	}
	if u.ActivityLevel == 0 {
		return false
	}
	if u.WeightGoal == "" {
		return false
	}
	if u.Email == "" {
		return false
	}
	return true
}
