package role

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
)

func UpdateRole(ctx context.Context, req *model.UpdateRoleRequest) (*model.Role, error) {
	role, err := GetRole(ctx, &model.GetRoleRequest{ID: req.ID})
	if err != nil {
		return nil, err
	}

	role.Name = req.Name

	if err := db.DB.WithContext(ctx).Save(role).Error; err != nil {
		log.Errorf("Failed to update role: %s", err)
		return nil, errors.Internal()
	}

	return role, nil
}
