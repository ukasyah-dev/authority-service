package role

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

func GetRole(ctx context.Context, req *model.GetRoleRequest) (*model.Role, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	role := new(model.Role)

	err := db.DB.WithContext(ctx).
		Where("id = ?", req.ID).
		Take(role).Error
	if err != nil {
		if e.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound()
		}

		log.Errorf("Failed to get role: %s", err)
		return nil, errors.Internal()
	}

	return role, nil
}
