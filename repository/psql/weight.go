package psql

import (
	"fmt"
	"time"

	"github.com/mbvisti/integration-test-in-go/entity"
	"github.com/pkg/errors"
)

const weightTblName = "weight"

func (s Storage) CreateWeightEntry(newWeight entity.Weight,
	createdAt time.Time) error {
	newWeightStatement := fmt.Sprintf(`
		INSERT INTO %s 
			(created_at, weight, user_id, bmr, daily_caloric_intake) 
		VALUES 
			($1, $2, $3, $4)
		`, weightTblName)
	_, err := s.db.Exec(newWeightStatement, createdAt, newWeight.Weight,
		newWeight.UserID, newWeight.BMR, newWeight.DailyCaloricIntake)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
