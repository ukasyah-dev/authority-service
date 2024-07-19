package action

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
)

func DeleteAction(ctx context.Context, req *model.DeleteActionRequest) (*model.Action, error) {
	action, err := GetAction(ctx, &model.GetActionRequest{ID: req.ID})
	if err != nil {
		return nil, err
	}

	if err := db.DB.WithContext(ctx).Delete(action).Error; err != nil {
		log.Errorf("Failed to delete action: %s", err)
		return nil, errors.Internal()
	}

	return action, nil
}
