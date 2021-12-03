package psql

import (
	"fmt"
	"time"

	"github.com/mbvisti/integration-test-in-go/entity"
	"github.com/pkg/errors"
)

const userTblName = "user"

func (s Storage) CreateUser(newUser entity.User, createdAt time.Time) error {
	insertStmt := fmt.Sprintf(`
		INSERT INTO 
			%s (created_at, name, age, height, sex, activity_level, email, weight_goal) 
		VALUES 
			($1, $2, $3, $4, $5, $6, $7);
	`, userTblName)

	_, err := s.db.Exec(insertStmt, createdAt, newUser.Name, newUser.Age,
		newUser.Height, newUser.Sex, newUser.ActivityLevel, newUser.Email,
		newUser.WeightGoal,
	)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
