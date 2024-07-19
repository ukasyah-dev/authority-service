package action_test

import (
	"net/http"
	"testing"

	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"github.com/ukasyah-dev/authority-service/tests"
	"github.com/ukasyah-dev/common/rest/testkit"
)

func TestGetAction_Success(t *testing.T) {
	testkit.New(tests.RESTServer).
		Get("/actions/"+tests.Data.Actions[2].ID).
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[2]).
		Expect(t).
		Status(http.StatusOK).
		Assert(jsonpath.Equal("$.id", tests.Data.Actions[2].ID)).
		Assert(jsonpath.Equal("$.name", tests.Data.Actions[2].Name)).
		End()
}
