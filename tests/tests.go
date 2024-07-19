package tests

import (
	"context"
	"os"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ukasyah-dev/authority-service/controller/action"
	"github.com/ukasyah-dev/authority-service/controller/permission"
	"github.com/ukasyah-dev/authority-service/controller/role"
	"github.com/ukasyah-dev/authority-service/controller/team"
	"github.com/ukasyah-dev/authority-service/controller/team_member"
	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/authority-service/rest"
	"github.com/ukasyah-dev/authority-service/rpc"
	"github.com/ukasyah-dev/common/amqp"
	commonAuth "github.com/ukasyah-dev/common/auth"
	dt "github.com/ukasyah-dev/common/db/testkit"
	gt "github.com/ukasyah-dev/common/grpc/testkit"
	"github.com/ukasyah-dev/common/hash"
	"github.com/ukasyah-dev/common/id"
	restServer "github.com/ukasyah-dev/common/rest/server"
	identityModel "github.com/ukasyah-dev/identity-service/model"
	pb "github.com/ukasyah-dev/pb/authority"
	"google.golang.org/grpc"
)

var AuthorityClient pb.AuthorityClient
var authorityClientConn *grpc.ClientConn
var closeAuthorityClientConn func()

var Data struct {
	AccessTokens []string
	Actions      []*model.Action
	Permissions  []*model.Permission
	Roles        []*model.Role
	Teams        []*model.Team
	TeamMembers  []*model.TeamMember
	Users        []*identityModel.User
}

var RESTServer *restServer.Server

func Setup() {
	faker.SetGenerateUniqueValues(true)

	amqp.Open(os.Getenv("AMQP_URL"))
	amqp.DeclareQueues("user-mutation")
	dt.CreateTestDB()
	db.Open()

	// Setup authority client
	grpcServer := grpc.NewServer()
	pb.RegisterAuthorityServer(grpcServer, &rpc.Server{})
	authorityClientConn, closeAuthorityClientConn = gt.NewClientConn(grpcServer)
	AuthorityClient = pb.NewAuthorityClient(authorityClientConn)

	RESTServer = rest.NewServer(AuthorityClient)

	privateKey, err := commonAuth.ParsePrivateKeyFromBase64(os.Getenv("BASE64_JWT_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	for i := 0; i <= 4; i++ {
		// Actions
		a, _ := action.CreateAction(ctx, &model.CreateActionRequest{
			ID:   id.New(),
			Name: faker.Name(),
		})
		Data.Actions = append(Data.Actions, a)

		// Roles
		r, _ := role.CreateRole(ctx, &model.CreateRoleRequest{
			ID:   id.New(),
			Name: faker.Name(),
		})
		Data.Roles = append(Data.Roles, r)

		// Permissions
		p, _ := permission.CreatePermission(ctx, &model.CreatePermissionRequest{
			ActionID: a.ID,
			RoleID:   r.ID,
		})
		Data.Permissions = append(Data.Permissions, p)

		// Users
		u := &identityModel.User{
			ID:         id.New(),
			Name:       faker.Name(),
			Email:      faker.Email(),
			Password:   hash.Generate(faker.Password()),
			Status:     "active",
			SuperAdmin: i == 0,
		}
		if err := db.DB.Create(u).Error; err != nil {
			panic(err)
		}
		Data.Users = append(Data.Users, u)

		// Access tokens
		accessToken, _ := commonAuth.GenerateAccessToken(privateKey, commonAuth.Claims{
			UserID:     u.ID,
			SuperAdmin: u.SuperAdmin,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			},
		})
		Data.AccessTokens = append(Data.AccessTokens, accessToken)

		// Teams
		t, _ := team.CreateTeam(ctx, &model.CreateTeamRequest{
			Name: faker.Name(),
		})
		Data.Teams = append(Data.Teams, t)

		// Team members
		tm, _ := team_member.CreateTeamMember(ctx, &model.CreateTeamMemberRequest{
			RoleID: "admin",
			TeamID: t.ID,
			UserID: u.ID,
		})
		Data.TeamMembers = append(Data.TeamMembers, tm)
	}
}

func Teardown() {
	closeAuthorityClientConn()
	amqp.Close()
	db.Close()
	dt.DestroyTestDB()
}
