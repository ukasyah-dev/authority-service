package rest

import (
	"net/http"
	"os"

	"github.com/swaggest/openapi-go/openapi31"
	"github.com/ukasyah-dev/authority-service/controller/invitation"
	"github.com/ukasyah-dev/authority-service/controller/role"
	"github.com/ukasyah-dev/authority-service/controller/team"
	commonAuth "github.com/ukasyah-dev/common/auth"
	"github.com/ukasyah-dev/common/rest/handler"
	"github.com/ukasyah-dev/common/rest/server"
)

var Server *server.Server

func init() {
	description := "Team management and authorization."
	spec := openapi31.Spec{
		Openapi: "3.1.0",
		Info: openapi31.Info{
			Title:       "Authority Service",
			Version:     "0.0.1",
			Description: &description,
		},
		Servers: []openapi31.Server{
			{URL: os.Getenv("OPENAPI_SERVER_URL")},
		},
	}

	// Parse JWT public key
	jwtPublicKey, err := commonAuth.ParsePublicKeyFromBase64(os.Getenv("BASE64_JWT_PUBLIC_KEY"))
	if err != nil {
		panic(err)
	}

	// Create new server
	Server = server.New(server.Config{
		OpenAPI:      server.OpenAPI{Spec: &spec},
		JWTPublicKey: jwtPublicKey,
	})

	handler.AddHealthCheck(Server)

	// Role
	handler.Add(Server, http.MethodPost, "/roles", role.CreateRole, handler.Config{
		Summary:     "Create role",
		Description: "Create role",
		Tags:        []string{"Role"},
		SuperAdmin:  true,
	})
	handler.Add(Server, http.MethodGet, "/roles", role.GetRoles, handler.Config{
		Summary:      "Get roles",
		Description:  "Get roles",
		Tags:         []string{"Role"},
		Authenticate: true,
	})
	handler.Add(Server, http.MethodGet, "/roles/:roleId", role.GetRole, handler.Config{
		Summary:      "Get role",
		Description:  "Get role",
		Tags:         []string{"Role"},
		Authenticate: true,
	})
	handler.Add(Server, http.MethodPatch, "/roles/:roleId", role.UpdateRole, handler.Config{
		Summary:     "Update role",
		Description: "Update role",
		Tags:        []string{"Role"},
		SuperAdmin:  true,
	})
	handler.Add(Server, http.MethodDelete, "/roles/:roleId", role.DeleteRole, handler.Config{
		Summary:     "Delete role",
		Description: "Delete role",
		Tags:        []string{"Role"},
		SuperAdmin:  true,
	})

	// Team
	handler.Add(Server, http.MethodPost, "/teams", team.CreateTeam, handler.Config{
		Summary:      "Create team",
		Description:  "Create team",
		Tags:         []string{"Team"},
		Authenticate: true,
	})
	handler.Add(Server, http.MethodGet, "/teams", team.GetTeams, handler.Config{
		Summary:      "Get teams",
		Description:  "Get teams",
		Tags:         []string{"Team"},
		Authenticate: true,
	})
	handler.Add(Server, http.MethodGet, "/teams/:teamId", team.GetTeam, handler.Config{
		Summary:      "Get team",
		Description:  "Get team",
		Tags:         []string{"Team"},
		Authenticate: true,
	})
	handler.Add(Server, http.MethodPatch, "/teams/:teamId", team.UpdateTeam, handler.Config{
		Summary:      "Update team",
		Description:  "Update team",
		Tags:         []string{"Team"},
		Authenticate: true,
	})
	handler.Add(Server, http.MethodDelete, "/teams/:teamId", team.DeleteTeam, handler.Config{
		Summary:      "Delete team",
		Description:  "Delete team",
		Tags:         []string{"Team"},
		Authenticate: true,
	})

	// Invitation
	handler.Add(Server, http.MethodPost, "/teams/:teamId/invitations", invitation.CreateInvitation, handler.Config{
		Summary:      "Create invitation",
		Description:  "Create invitation",
		Tags:         []string{"Invitation"},
		Authenticate: true,
	})

}
