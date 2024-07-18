package model

import (
	"time"

	commonModel "github.com/ukasyah-dev/common/model"
)

type Team struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type CreateTeamRequest struct {
	Name string `json:"name" validate:"required,lte=30"`
}

type GetTeamsRequest struct {
	commonModel.PaginationRequest
}

type GetTeamsResponse struct {
	commonModel.PaginationResponse
	Data []*Team `json:"data"`
}

type GetTeamRequest struct {
	ID string `params:"teamId" path:"teamId" validate:"required"`
}

type UpdateTeamRequest struct {
	ID   string `params:"teamId" path:"teamId" validate:"required"`
	Name string `json:"name" validate:"omitempty,lte=30"`
}

type DeleteTeamRequest struct {
	ID string `params:"teamId" path:"teamId" validate:"required"`
}
