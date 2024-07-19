package role

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
	"github.com/ukasyah-dev/common/validator"
)

func CreateRole(ctx context.Context, req *model.CreateRoleRequest) (*model.Role, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	role := &model.Role{
		ID:   req.ID,
		Name: req.Name,
	}

	if err := db.DB.WithContext(ctx).Create(role).Error; err != nil {
		log.Errorf("Failed to create role: %s", err)
		return nil, errors.Internal()
	}

	return role, nil
}
