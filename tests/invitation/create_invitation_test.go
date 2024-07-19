package invitation_test

import (
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/authority-service/tests"
	"github.com/ukasyah-dev/common/rest/testkit"
)

func TestCreateInvitation_Success(t *testing.T) {
	testkit.New(tests.RESTServer).
		Post("/teams/"+tests.Data.Teams[0].ID+"/invitations").
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[0]).
		JSON(map[string]any{
			"email":  faker.Email(),
			"roleId": tests.Data.Roles[0].ID,
			"teamId": tests.Data.Teams[0].ID,
		}).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Present("$.id")).
		End()
}
