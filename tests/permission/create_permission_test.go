package permission_test

import (
	"net/http"
	"testing"

	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/authority-service/tests"
	"github.com/ukasyah-dev/common/rest/testkit"
)

func TestCreatePermission_Success(t *testing.T) {
	testkit.New(tests.RESTServer).
		Post("/permissions").
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[0]).
		JSON(map[string]any{
			"actionId": tests.Data.Actions[0].ID,
			"roleId":   tests.Data.Roles[1].ID,
		}).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Present("$.id")).
		End()
}
