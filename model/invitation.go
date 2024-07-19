package model

import "time"

type Invitation struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Email     string    `json:"email"`
	RoleID    string    `json:"roleId"`
	Role      *Role     `gorm:"constraint:OnDelete:CASCADE" json:"role,omitempty"`
	TeamID    string    `gorm:"index:idx_unique_team_member,unique" json:"teamId"`
	Team      *Team     `gorm:"constraint:OnDelete:CASCADE" json:"team,omitempty"`
	Token     string    `gorm:"unique" json:"token"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateInvitationRequest struct {
	Email  string `json:"email" validate:"required,email,lte=30" example:"user@example.com"`
	RoleID string `json:"roleId" validate:"required"`
	TeamID string `params:"teamId" path:"teamId" validate:"required"`
}
