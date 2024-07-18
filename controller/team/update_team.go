package team

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
)

func UpdateTeam(ctx context.Context, req *model.UpdateTeamRequest) (*model.Team, error) {
	t, err := GetTeam(ctx, &model.GetTeamRequest{ID: req.ID})
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		t.Name = req.Name
	}

	if err := db.DB.WithContext(ctx).Save(t).Error; err != nil {
		log.Errorf("Failed to update team: %s", err)
		return nil, errors.Internal()
	}

	return t, nil
}
