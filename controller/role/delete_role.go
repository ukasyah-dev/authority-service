package role

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
)

func DeleteRole(ctx context.Context, req *model.DeleteRoleRequest) (*model.Role, error) {
	role, err := GetRole(ctx, &model.GetRoleRequest{ID: req.ID})
	if err != nil {
		return nil, err
	}

	if err := db.DB.WithContext(ctx).Delete(role).Error; err != nil {
		log.Errorf("Failed to delete role: %s", err)
		return nil, errors.Internal()
	}

	return role, nil
}
