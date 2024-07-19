package team_test

import (
	"net/http"
	"testing"

	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/authority-service/tests"
	"github.com/ukasyah-dev/common/rest/testkit"
)

func TestGetTeams_Success(t *testing.T) {
	testkit.New(tests.RESTServer).
		Get("/teams").
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[1]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.GreaterThan("$.data", 0)).
		Assert(jsonpath.Present("$.data[0].id")).
		End()
}
