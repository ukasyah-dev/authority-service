package role_test

import (
	"net/http"
	"testing"

	"github.com/ukasyah-dev/authority-service/tests"
	"github.com/ukasyah-dev/common/rest/testkit"
)

func TestDeleteRole_Success(t *testing.T) {
	testkit.New(tests.RESTServer).
		Delete("/roles/"+tests.Data.Roles[3].ID).
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[0]).
		Expect(t).
		Status(http.StatusOK).
		End()
}
