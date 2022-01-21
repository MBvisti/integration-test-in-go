package http_test

import (
	"bytes"
	"encoding/json"
	h "net/http"

	"github.com/mbvisti/integration-test-in-go/server/http"
)

func (suite *HttpTestSuite) TestIntegration_UserHandler_CreateUser() {
	req := http.NewUserRequest{
		Name:          "Test user",
		Sex:           "male",
		WeightGoal:    "80",
		Email:         "test@gmail.com",
		Age:           99,
		Height:        185,
		ActivityLevel: 1,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(req)
	suite.NoError(err)

	rq, err := h.NewRequest(h.MethodPost, "/api/user", &buf)
	suite.NoError(err)

	rq.Header.Add("Content-Type", "application/json")

	res, err := suite.TestRouter.Test(rq, -1)
	suite.NoError(err)

	suite.EqualValues(200, res.StatusCode)

	// query the database to verify that a user was created based on the request
	// we sent
	newUser, err := suite.TestStorage.GetUserFromEmail(req.Email)
	suite.NoError(err)

	suite.EqualValues(newUser.Height, req.Height)
	suite.EqualValues(newUser.Name, req.Name)
	suite.EqualValues(newUser.Sex, req.Sex)
	suite.EqualValues(newUser.WeightGoal, req.WeightGoal)
	suite.EqualValues(newUser.Email, req.Email)
	suite.EqualValues(newUser.Age, req.Age)
	suite.EqualValues(newUser.ActivityLevel, req.ActivityLevel)
}
