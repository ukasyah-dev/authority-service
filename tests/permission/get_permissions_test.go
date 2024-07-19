package permission_test

import (
	"net/http"
	"testing"

	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/authority-service/rest"
	"github.com/ukasyah-dev/authority-service/tests"
	"github.com/ukasyah-dev/common/rest/testkit"
)

func TestGetPermissions_Success(t *testing.T) {
	testkit.New(rest.Server).
		Get("/permissions").
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[1]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.GreaterThan("$.data", 0)).
		Assert(jsonpath.Present("$.data[0].id")).
		End()
}
