package http_test

import (
	"bytes"
	"encoding/json"
	h "net/http"

	"github.com/mbvisti/integration-test-in-go/server/http"
)

func (suite *HttpTestSuite) TestServer_WeightHandler_CreateWeight() {
	req := http.NewWeightRequest{
		Weight:             108,
		UserId:             1,
		BMR:                3103,
		DailyCaloricIntake: 2500,
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(req)
	suite.NoError(err)

	rq, err := h.NewRequest(h.MethodPost, "/api/weight", &buf)
	suite.NoError(err)

	rq.Header.Add("Content-Type", "application/json")

	res, err := suite.TestRouter.Test(rq, -1)
	suite.NoError(err)

	// a little bit lazy here, I know
	suite.EqualValues(200, res.StatusCode)
}
