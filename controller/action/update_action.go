package action

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
)

func UpdateAction(ctx context.Context, req *model.UpdateActionRequest) (*model.Action, error) {
	action, err := GetAction(ctx, &model.GetActionRequest{ID: req.ID})
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		action.Name = req.Name
	}

	if err := db.DB.WithContext(ctx).Save(action).Error; err != nil {
		log.Errorf("Failed to update action: %s", err)
		return nil, errors.Internal()
	}

	return action, nil
}
