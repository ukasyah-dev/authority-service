package rest

import (
	"net/http"
	"os"

	"github.com/swaggest/openapi-go/openapi31"
	"github.com/ukasyah-dev/authority-service/controller/action"
	"github.com/ukasyah-dev/authority-service/controller/invitation"
	"github.com/ukasyah-dev/authority-service/controller/permission"
	"github.com/ukasyah-dev/authority-service/controller/role"
	"github.com/ukasyah-dev/authority-service/controller/team"
	"github.com/ukasyah-dev/authority-service/controller/team_member"
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

	// Action
	handler.Add(Server, http.MethodPost, "/actions", action.CreateAction, handler.Config{
		Summary:     "Create action",
		Description: "Create action",
		Tags:        []string{"Action"},
		SuperAdmin:  true,
	})
	handler.Add(Server, http.MethodGet, "/actions", action.GetActions, handler.Config{
		Summary:      "Get Actions",
		Description:  "Get Actions",
		Tags:         []string{"Action"},
		Authenticate: true,
	})
	handler.Add(Server, http.MethodGet, "/actions/:actionId", action.GetAction, handler.Config{
		Summary:      "Get action",
		Description:  "Get action",
		Tags:         []string{"Action"},
		Authenticate: true,
	})
	handler.Add(Server, http.MethodPatch, "/actions/:actionId", action.UpdateAction, handler.Config{
		Summary:     "Update action",
		Description: "Update action",
		Tags:        []string{"Action"},
		SuperAdmin:  true,
	})
	handler.Add(Server, http.MethodDelete, "/actions/:actionId", action.DeleteAction, handler.Config{
		Summary:     "Delete action",
		Description: "Delete action",
		Tags:        []string{"Action"},
		SuperAdmin:  true,
	})

	// Permission
	handler.Add(Server, http.MethodPost, "/permissions", permission.CreatePermission, handler.Config{
		Summary:     "Create permission",
		Description: "Create permission",
		Tags:        []string{"Permission"},
		SuperAdmin:  true,
	})
	handler.Add(Server, http.MethodGet, "/permissions", permission.GetPermissions, handler.Config{
		Summary:      "Get permissions",
		Description:  "Get permissions",
		Tags:         []string{"Permission"},
		Authenticate: true,
	})
	handler.Add(Server, http.MethodGet, "/permissions/:permissionId", permission.GetPermission, handler.Config{
		Summary:      "Get permission",
		Description:  "Get permission",
		Tags:         []string{"Permission"},
		Authenticate: true,
	})
	handler.Add(Server, http.MethodDelete, "/permissions/:permissionId", permission.DeletePermission, handler.Config{
		Summary:     "Delete permission",
		Description: "Delete permission",
		Tags:        []string{"Permission"},
		SuperAdmin:  true,
	})

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
		Summary:     "Get team",
		Description: "Get team",
		Tags:        []string{"Team"},
		Permission:  "read-team",
	})
	handler.Add(Server, http.MethodPatch, "/teams/:teamId", team.UpdateTeam, handler.Config{
		Summary:     "Update team",
		Description: "Update team",
		Tags:        []string{"Team"},
		Permission:  "write-team",
	})
	handler.Add(Server, http.MethodDelete, "/teams/:teamId", team.DeleteTeam, handler.Config{
		Summary:      "Delete team",
		Description:  "Delete team",
		Tags:         []string{"Team"},
		Authenticate: true,
		Permission:   "write-team",
	})

	// Invitation
	handler.Add(Server, http.MethodPost, "/teams/:teamId/invitations", invitation.CreateInvitation, handler.Config{
		Summary:      "Create invitation",
		Description:  "Create invitation",
		Tags:         []string{"Invitation"},
		Authenticate: true,
	})

	// Team member
	handler.Add(Server, http.MethodGet, "/teams/:teamId/members", team_member.GetTeamMembers, handler.Config{
		Summary:     "Get team members",
		Description: "Get team members",
		Tags:        []string{"Team member"},
		Permission:  "read-team",
	})
	handler.Add(Server, http.MethodGet, "/teams/:teamId/members/:teamMemberId", team_member.GetTeamMember, handler.Config{
		Summary:     "Get team member",
		Description: "Get team member",
		Tags:        []string{"Team member"},
		Permission:  "read-team",
	})
	handler.Add(Server, http.MethodPatch, "/teams/:teamId/members/:teamMemberId", team_member.UpdateTeamMember, handler.Config{
		Summary:     "Update team member",
		Description: "Update team member",
		Tags:        []string{"Team member"},
		Permission:  "write-team",
	})
	handler.Add(Server, http.MethodDelete, "/teams/:teamId/members/:teamMemberId", team_member.DeleteTeamMember, handler.Config{
		Summary:     "Delete team member",
		Description: "Delete team member",
		Tags:        []string{"Team member"},
		Permission:  "write-team",
	})
}
