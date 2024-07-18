package team_test

import (
	"os"
	"testing"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/tests"
	"github.com/ukasyah-dev/common/amqp"
	dt "github.com/ukasyah-dev/common/db/testkit"
)

func TestMain(m *testing.M) {
	amqp.Open(os.Getenv("AMQP_URL"))
	dt.CreateTestDB()
	db.Open()
	tests.Setup()

	code := m.Run()

	amqp.Close()
	db.Close()
	dt.DestroyTestDB()

	os.Exit(code)
}
