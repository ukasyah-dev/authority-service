package model

import (
	"time"

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
