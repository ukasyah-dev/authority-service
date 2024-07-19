package permission

import (
	"context"
	e "errors"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
	"github.com/ukasyah-dev/common/validator"
	"gorm.io/gorm"
)

func GetPermission(ctx context.Context, req *model.GetPermissionRequest) (*model.Permission, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	p := new(model.Permission)

	err := db.DB.WithContext(ctx).
		Where("id = ?", req.ID).
		Take(p).Error
	if err != nil {
		if e.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound()
		}

		log.Errorf("Failed to get permission: %s", err)
		return nil, errors.Internal()
	}

	return p, nil
}
