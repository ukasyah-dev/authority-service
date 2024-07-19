package permission

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/id"
	"github.com/ukasyah-dev/common/log"
	"github.com/ukasyah-dev/common/validator"
)

func CreatePermission(ctx context.Context, req *model.CreatePermissionRequest) (*model.Permission, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	p := &model.Permission{
		ID:       id.New(),
		ActionID: req.ActionID,
		RoleID:   req.RoleID,
	}

	if err := db.DB.WithContext(ctx).Create(p).Error; err != nil {
		log.Errorf("Failed to create permission: %s", err)
		return nil, errors.Internal()
	}

	return p, nil
}
