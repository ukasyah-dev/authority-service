package team_member

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
)

func UpdateTeamMember(ctx context.Context, req *model.UpdateTeamMemberRequest) (*model.TeamMember, error) {
	t, err := GetTeamMember(ctx, &model.GetTeamMemberRequest{
		TeamID:       req.TeamID,
		TeamMemberID: req.TeamMemberID,
	})
	if err != nil {
		return nil, err
	}

	t.RoleID = req.RoleID

	if err := db.DB.WithContext(ctx).Save(t).Error; err != nil {
		log.Errorf("Failed to update team member: %s", err)
		return nil, errors.Internal()
	}

	return t, nil
}
