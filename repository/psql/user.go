package psql

import (
	"github.com/mbvisti/integration-test-in-go/entity"
	"github.com/pkg/errors"
)

const userTblName = "user"

func (s Storage) CreateUser(newUser entity.User) error {
	insertStmt := `INSERT INTO users (name, age, height, sex, activity_level, email, 
		weight_goal) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := s.db.Exec(insertStmt, newUser.Name, newUser.Age,
		newUser.Height, newUser.Sex, newUser.ActivityLevel, newUser.Email,
		newUser.WeightGoal,
	)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s Storage) GetUserFromEmail(email string) (*entity.User, error) {
	getStmt := `
		SELECT id, name, age, height, sex, activity_level, email, weight_goal
		FROM users
		WHERE users.email=$1
	`

	var user entity.User
	err := s.db.QueryRow(getStmt, email).Scan(
		&user.ID, &user.Name, &user.Age, &user.Height, &user.Sex, &user.ActivityLevel,
		&user.Email, &user.WeightGoal,
	)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &user, nil
}
