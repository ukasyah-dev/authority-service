package team_test

import (
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/ukasyah-dev/authority-service/rest"
	"github.com/ukasyah-dev/authority-service/tests"
	"github.com/ukasyah-dev/common/rest/testkit"
)

func TestDeleteTeam_Success(t *testing.T) {
	testkit.New(rest.Server).
		Delete("/teams/"+tests.Data.Teams[3].ID).
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[3]).
		JSON(map[string]any{
			"name": faker.Name(),
		}).
		Expect(t).
		Status(http.StatusOK).
		End()
}
