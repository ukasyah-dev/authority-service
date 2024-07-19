package main

import (
	"context"
	"os"

	"github.com/appleboy/graceful"
	"github.com/caitlinelfring/go-env-default"
	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/rest"
	"github.com/ukasyah-dev/common/amqp"
	identityModel "github.com/ukasyah-dev/identity-service/model"
	pb "github.com/ukasyah-dev/pb/authority"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var port = env.GetIntDefault("PORT", 3000)

func init() {
	amqp.Open(os.Getenv("AMQP_URL"))
	amqp.DeclareQueues("user-mutation")
	db.Open()
}

func main() {
	// Create authority client
	addr := "localhost:" + env.GetDefault("GRPC_PORT", "4000")
	authorityClientConn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	authorityClient := pb.NewAuthorityClient(authorityClientConn)

	s := rest.NewServer(authorityClient)

	m := graceful.NewManager()

	m.AddRunningJob(func(ctx context.Context) error {
		return amqp.ConsumeMutation(ctx, db.DB, "user-mutation", "authority-service", &identityModel.User{})
	})

	m.AddRunningJob(func(ctx context.Context) error {
		return s.Start(port)
	})

	m.AddShutdownJob(func() error {
		return s.Shutdown()
	})

	m.AddShutdownJob(func() error {
		return amqp.Close()
	})

	m.AddShutdownJob(func() error {
		return db.Close()
	})

	<-m.Done()
}
