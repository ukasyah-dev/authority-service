package permission

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
)

func DeletePermission(ctx context.Context, req *model.DeletePermissionRequest) (*model.Permission, error) {
	p, err := GetPermission(ctx, &model.GetPermissionRequest{ID: req.ID})
	if err != nil {
		return nil, err
	}

	if err := db.DB.WithContext(ctx).Delete(p).Error; err != nil {
		log.Errorf("Failed to delete permission: %s", err)
		return nil, errors.Internal()
	}

	return p, nil
}
