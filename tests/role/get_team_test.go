package team_test

import (
	"net/http"
	"testing"

	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/authority-service/rest"
	"github.com/ukasyah-dev/authority-service/tests"
	"github.com/ukasyah-dev/common/rest/testkit"
)

func TestGetTeam_Success(t *testing.T) {
	testkit.New(rest.Server).
		Get("/teams/"+tests.Data.Teams[2].ID).
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[2]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Equal("$.id", tests.Data.Teams[2].ID)).
		Assert(jsonpath.Equal("$.name", tests.Data.Teams[2].Name)).
		End()
}
