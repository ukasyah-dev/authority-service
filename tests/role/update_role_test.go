package role_test

import (
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/authority-service/rest"
	"github.com/ukasyah-dev/authority-service/tests"
	"github.com/ukasyah-dev/common/rest/testkit"
)

func TestUpdateRole_Success(t *testing.T) {
	data := map[string]any{
		"name": faker.Name(),
	}

	testkit.New(rest.Server).
		Patch("/roles/"+tests.Data.Roles[0].ID).
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[0]).
		JSON(data).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Equal("$.id", tests.Data.Roles[0].ID)).
		Assert(jsonpath.Equal("$.name", data["name"])).
		End()
}
