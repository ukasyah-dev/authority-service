package permission_test

import (
	"net/http"
	"testing"

	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/authority-service/rest"
	"github.com/ukasyah-dev/authority-service/tests"
	"github.com/ukasyah-dev/common/rest/testkit"
)

func TestGetPermission_Success(t *testing.T) {
	testkit.New(rest.Server).
		Get("/permissions/"+tests.Data.Permissions[2].ID).
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[2]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Equal("$.id", tests.Data.Permissions[2].ID)).
		End()
}
