package team_member_test

import (
	"os"
	"testing"

	"github.com/ukasyah-dev/authority-service/tests"
)

func TestMain(m *testing.M) {
	tests.Setup()
	defer tests.Teardown()
	code := m.Run()
	os.Exit(code)
}
