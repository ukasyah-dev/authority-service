package team

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

func GetTeam(ctx context.Context, req *model.GetTeamRequest) (*model.Team, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	team := &model.Team{
		ID: req.ID,
	}

	if err := db.DB.WithContext(ctx).Take(team).Error; err != nil {
		if e.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound()
		}

		log.Errorf("Failed to get team: %s", err)
		return nil, errors.Internal()
	}

	return team, nil
}
