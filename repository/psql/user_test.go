package psql_test

import (
	_ "github.com/lib/pq"
	"github.com/mbvisti/integration-test-in-go/entity"
)

// testing the happy path only - to improve upon these tests, we could consider
// using a table test
func (suite *RepositoryTestSuite) TestStorage_CreateUser() {
	// run the test
	newUser, err := entity.NewUser(
		"Jon Snow", "male", "90", "theyoungwolf@stark.com", 16, 182, 1)
	suite.NoError(err)

	err = suite.TestStorage.CreateUser(*newUser)
	// assert there is no err
	suite.NoError(err)

	queryResult := entity.User{}
	err = suite.TestDb.QueryRow("SELECT id, name, email FROM users WHERE email=$1",
		"theyoungwolf@stark.com").Scan(
		&queryResult.ID, &queryResult.Name, &queryResult.Email,
	)
	suite.NoError(err)

	suite.EqualValues(queryResult.Name, newUser.Name)
	suite.EqualValues(queryResult.Email, newUser.Email)
}
