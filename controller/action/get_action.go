package action

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

func GetAction(ctx context.Context, req *model.GetActionRequest) (*model.Action, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	action := new(model.Action)

	err := db.DB.WithContext(ctx).
		Where("id = ?", req.ID).
		Take(action).Error
	if err != nil {
		if e.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound()
		}

		log.Errorf("Failed to get action: %s", err)
		return nil, errors.Internal()
	}

	return action, nil
}
