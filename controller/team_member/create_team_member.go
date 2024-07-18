package team_member

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
	"github.com/ukasyah-dev/common/validator"
)

func CreateTeamMember(ctx context.Context, req *model.CreateTeamMemberRequest) (*model.TeamMember, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	tm := &model.TeamMember{
		RoleID: req.RoleID,
		TeamID: req.TeamID,
		UserID: req.UserID,
	}

	if err := db.DB.WithContext(ctx).Create(tm).Error; err != nil {
		log.Errorf("Failed to create team member: %s", err)
		return nil, errors.Internal()
	}

	return tm, nil
}
