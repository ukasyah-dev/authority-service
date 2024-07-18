package team_member

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

func GetTeamMember(ctx context.Context, req *model.GetTeamMemberRequest) (*model.TeamMember, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	tm := new(model.TeamMember)

	err := db.DB.WithContext(ctx).
		Where("team_id = ?", req.TeamID).
		Where("user_id = ?", req.UserID).
		Take(tm).Error
	if err != nil {
		if e.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.NotFound()
		}

		log.Errorf("Failed to get team member: %s", err)
		return nil, errors.Internal()
	}

	return tm, nil
}
