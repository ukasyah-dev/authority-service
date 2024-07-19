package action

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
	"github.com/ukasyah-dev/common/validator"
)

func CreateAction(ctx context.Context, req *model.CreateActionRequest) (*model.Action, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	action := &model.Action{
		ID:   req.ID,
		Name: req.Name,
	}

	if err := db.DB.WithContext(ctx).Create(action).Error; err != nil {
		log.Errorf("Failed to create action: %s", err)
		return nil, errors.Internal()
	}

	return action, nil
}
