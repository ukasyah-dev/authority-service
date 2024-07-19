package model

import commonModel "github.com/ukasyah-dev/common/model"

type Permission struct {
	ID       string  `gorm:"primaryKey" json:"id"`
	ActionID string  `gorm:"index:idx_unique_permission,unique" json:"actionId"`
	Action   *Action `gorm:"constraint:OnDelete:CASCADE" json:"action,omitempty"`
	RoleID   string  `gorm:"index:idx_unique_permission,unique" json:"roleId"`
	Role     *Role   `gorm:"constraint:OnDelete:CASCADE" json:"role,omitempty"`
}

type CreatePermissionRequest struct {
	ActionID string `json:"actionId" validate:"required"`
	RoleID   string `json:"roleId" validate:"required"`
}

type CheckPermissionRequest struct {
	ActionID string `json:"actionId" validate:"required"`
	UserID   string `json:"userId" validate:"required"`
}

type CheckPermissionResponse struct {
	Allowed bool `json:"allowed"`
}

type GetPermissionsRequest struct {
	commonModel.PaginationRequest
}

type GetPermissionsResponse struct {
	commonModel.PaginationResponse
	Data []*Permission `json:"data"`
}

type GetPermissionRequest struct {
	ID string `params:"permissionId" path:"permissionId" validate:"required"`
}

type DeletePermissionRequest struct {
	ID string `params:"permissionId" path:"permissionId" validate:"required"`
}
