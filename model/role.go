package model

type Role struct {
	ID   string `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type CreateRoleRequest struct {
	ID   string `json:"id" validate:"required,lte=30"`
	Name string `json:"name" validate:"required,lte=40"`
}
