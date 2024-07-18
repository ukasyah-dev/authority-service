package team

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
)

func DeleteTeam(ctx context.Context, req *model.DeleteTeamRequest) (*model.Team, error) {
	t, err := GetTeam(ctx, &model.GetTeamRequest{
		ID: req.ID,
	})
	if err != nil {
		return nil, err
	}

	if err := db.DB.WithContext(ctx).Delete(t).Error; err != nil {
		log.Errorf("Failed to delete team: %s", err)
		return nil, errors.Internal()
	}

	return t, nil
}
