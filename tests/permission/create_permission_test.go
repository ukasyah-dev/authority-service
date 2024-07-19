package permission_test

import (
	"net/http"
	"testing"

	"github.com/go-faker/faker/v4"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/authority-service/rest"
	"github.com/ukasyah-dev/authority-service/tests"
	"github.com/ukasyah-dev/common/rest/testkit"
)

func TestCreatePermission_Success(t *testing.T) {
	testkit.New(rest.Server).
		Post("/permissions").
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[0]).
		JSON(map[string]any{
			"name": faker.Name(),
		}).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Present("$.id")).
		End()
}
