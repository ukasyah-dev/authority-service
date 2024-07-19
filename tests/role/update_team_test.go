package team_test

import (
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/authority-service/rest"
	"github.com/ukasyah-dev/authority-service/tests"
	"github.com/ukasyah-dev/common/id"
	"github.com/ukasyah-dev/common/rest/testkit"
)

func TestUpdateTeam_Success(t *testing.T) {
	data := map[string]any{
		"name": faker.Name(),
	}

	testkit.New(rest.Server).
		Patch("/teams/"+tests.Data.Teams[0].ID).
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[0]).
		JSON(data).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Equal("$.id", tests.Data.Teams[0].ID)).
		Assert(jsonpath.Equal("$.name", data["name"])).
		End()
}

func TestUpdateTeam_NotFound(t *testing.T) {
	testkit.New(rest.Server).
		Patch("/teams/"+id.New()).
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[0]).
		JSON(map[string]any{
			"name": faker.Name(),
		}).
		Expect(t).
		Status(http.StatusNotFound).
		End()
}
