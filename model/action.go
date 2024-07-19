package model

import commonModel "github.com/ukasyah-dev/common/model"

type Action struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type CreateActionRequest struct {
	ID   string `json:"id" validate:"required,lte=30"`
	Name string `json:"name" validate:"required,lte=30"`
}

type GetActionsRequest struct {
	commonModel.PaginationRequest
}

type GetActionsResponse struct {
	commonModel.PaginationResponse
	Data []*Action `json:"data"`
}

type GetActionRequest struct {
	ID string `params:"actionId" path:"actionId" validate:"required"`
}

type UpdateActionRequest struct {
	ID   string `params:"actionId" path:"actionId" validate:"required"`
	Name string `json:"name" validate:"required,lte=30"`
}

type DeleteActionRequest struct {
	ID string `params:"actionId" path:"actionId" validate:"required"`
}
