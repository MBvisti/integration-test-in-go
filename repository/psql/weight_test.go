package psql_test

import (
	"time"

	"github.com/mbvisti/integration-test-in-go/entity"
)

func (suite *RepositoryTestSuite) TestStorage_CreateWeightEntry() {
	newWeightEntry, err := entity.NewWeight(80, 2, 3210, 3300)
	suite.NoError(err)

	// to ensure consistency we could consider adding in a static date
	// i.e. time.Date(insert-fixed-date-here)
	// creationTime := time.Now()
	err = suite.TestStorage.CreateWeightEntry(*newWeightEntry, time.Now())
	// assert there is no err
	suite.NoError(err)

	queryResults := []entity.Weight{}
	rows, err := suite.TestDb.Query("SELECT id, created_at, updated_at, weight, user_id, bmr, daily_caloric_intake FROM weight WHERE user_id=$1", 2)
	suite.NoError(rows.Err())
	suite.NoError(err)

	for rows.Next() {
		weight := entity.Weight{}
		err := rows.Scan(&weight.ID, &weight.CreatedAt,
			&weight.UpdatedAt, &weight.Weight, &weight.UserID,
			&weight.BMR, &weight.DailyCaloricIntake,
		)
		suite.NoError(err)

		queryResults = append(queryResults, weight)
	}

	suite.True(len(queryResults) > 0)
	suite.EqualValues(queryResults[1].BMR, newWeightEntry.BMR)
}
