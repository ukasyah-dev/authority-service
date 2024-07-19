package action_test

import (
	"net/http"
	"testing"

	"github.com/ukasyah-dev/authority-service/rest"
	"github.com/ukasyah-dev/authority-service/tests"
	"github.com/ukasyah-dev/common/rest/testkit"
)

func TestDeleteAction_Success(t *testing.T) {
	testkit.New(rest.Server).
		Delete("/actions/"+tests.Data.Actions[3].ID).
		Header("Authorization", "Bearer "+tests.Data.AccessTokens[0]).
		Expect(t).
		Status(http.StatusOK).
		End()
}
