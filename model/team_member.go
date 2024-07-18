package model

import (
	"time"

	commonModel "github.com/ukasyah-dev/common/model"
	identityModel "github.com/ukasyah-dev/identity-service/model"
)

type TeamMember struct {
	ID        string              `gorm:"primaryKey" json:"id"`
	RoleID    string              `json:"roleId"`
	Role      *Role               `gorm:"constraint:OnDelete:CASCADE" json:"role,omitempty"`
	TeamID    string              `gorm:"index:idx_unique_team_member,unique" json:"teamId"`
	Team      *Team               `gorm:"constraint:OnDelete:CASCADE" json:"team,omitempty"`
	UserID    string              `gorm:"index:idx_unique_team_member,unique" json:"userId"`
	User      *identityModel.User `gorm:"constraint:OnDelete:CASCADE" json:"user,omitempty"`
	CreatedAt time.Time           `json:"createdAt"`
	UpdatedAt time.Time           `json:"updatedAt"`
}

type CreateTeamMemberRequest struct {
	TeamID string `json:"teamId"`
	UserID string `json:"userId"`
	RoleID string `json:"roleId"`
}

type GetTeamMembersRequest struct {
	commonModel.PaginationRequest
	TeamID string `params:"teamId" path:"teamId" validate:"required"`
}

type GetTeamMembersResponse struct {
	commonModel.PaginationResponse
	Data []*TeamMember `json:"data"`
}

type GetTeamMemberRequest struct {
	TeamID string `params:"teamId" path:"teamId" validate:"required"`
	UserID string `params:"userId" path:"userId" validate:"required"`
}

type UpdateTeamMemberRequest struct {
	TeamID string `params:"teamId" path:"teamId" validate:"required"`
	UserID string `params:"userId" path:"userId" validate:"required"`
	RoleID string `json:"roleId" validate:"required"`
}

type DeleteTeamMemberRequest struct {
	TeamID string `params:"teamId" path:"teamId" validate:"required"`
	UserID string `params:"userId" path:"userId" validate:"required"`
}
