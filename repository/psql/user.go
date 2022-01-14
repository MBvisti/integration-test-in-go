package psql

import (
	"log"

	"github.com/mbvisti/integration-test-in-go/entity"
	"github.com/pkg/errors"
)

const userTblName = "user"

func (s Storage) CreateUser(newUser entity.User) error {
	insertStmt := `INSERT INTO users (name, age, height, sex, activity_level, email, weight_goal) VALUES ($1, $2, $3, $4, $5, $6, $7)`

	_, err := s.db.Exec(insertStmt, newUser.Name, newUser.Age,
		newUser.Height, newUser.Sex, newUser.ActivityLevel, newUser.Email,
		newUser.WeightGoal,
	)
	if err != nil {
		log.Print(err)
		return errors.WithStack(err)
	}

	return nil
}
