package main

import (
	"context"
	"os"

	"github.com/appleboy/graceful"
	"github.com/caitlinelfring/go-env-default"
	"github.com/ukasyah-dev/authority-service/controller/user"
	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/rest"
	"github.com/ukasyah-dev/common/amqp"
)

var port = env.GetIntDefault("PORT", 3000)

func init() {
	amqp.Open(os.Getenv("AMQP_URL"))
	amqp.DeclareQueues("user-mutation")
	db.Open()
}

func main() {
	m := graceful.NewManager()

	m.AddRunningJob(func(ctx context.Context) error {
		return amqp.Consume(ctx, "user-mutation", "authority-service", user.HandleUserMutation)
	})

	m.AddRunningJob(func(ctx context.Context) error {
		return rest.Server.Start(port)
	})

	m.AddShutdownJob(func() error {
		return rest.Server.Shutdown()
	})

	m.AddShutdownJob(func() error {
		return amqp.Close()
	})

	m.AddShutdownJob(func() error {
		return db.Close()
	})

	<-m.Done()
}
