package tests

import (
	"context"
	"os"
	"time"

	"github.com/go-faker/faker/v4"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ukasyah-dev/authority-service/controller/team"
	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	commonAuth "github.com/ukasyah-dev/common/auth"
	"github.com/ukasyah-dev/common/hash"
	"github.com/ukasyah-dev/common/id"
	identityModel "github.com/ukasyah-dev/identity-service/model"
)

var Data struct {
	AccessTokens []string
	Teams        []*model.Team
	Users        []*identityModel.User
}

func Setup() {
	privateKey, err := commonAuth.ParsePrivateKeyFromBase64(os.Getenv("BASE64_JWT_PRIVATE_KEY"))
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	for i := 0; i <= 4; i++ {
		// Users
		user := &identityModel.User{
			ID:         id.New(),
			Name:       faker.Name(),
			Email:      faker.Email(),
			Password:   hash.Generate(faker.Password()),
			Status:     "active",
			SuperAdmin: i == 0,
		}
		if err := db.DB.Create(user).Error; err != nil {
			panic(err)
		}
		Data.Users = append(Data.Users, user)

		// Access tokens
		accessToken, _ := commonAuth.GenerateAccessToken(privateKey, commonAuth.Claims{
			UserID:     user.ID,
			SuperAdmin: user.SuperAdmin,
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
	}
}
