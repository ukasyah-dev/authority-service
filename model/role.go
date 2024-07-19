package model

import commonModel "github.com/ukasyah-dev/common/model"

type Role struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type CreateRoleRequest struct {
	ID   string `json:"id" validate:"required,lte=30"`
	Name string `json:"name" validate:"required,lte=30"`
}

type GetRolesRequest struct {
	commonModel.PaginationRequest
}

type GetRolesResponse struct {
	commonModel.PaginationResponse
	Data []*Role `json:"data"`
}

type GetRoleRequest struct {
	ID string `params:"roleId" path:"roleId" validate:"required"`
}

type UpdateRoleRequest struct {
	ID   string `params:"roleId" path:"roleId" validate:"required"`
	Name string `json:"name" validate:"required,lte=30"`
}

type DeleteRoleRequest struct {
	ID string `params:"roleId" path:"roleId" validate:"required"`
}
