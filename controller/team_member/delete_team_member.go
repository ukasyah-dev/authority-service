package team_member

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
)

func DeleteTeamMember(ctx context.Context, req *model.DeleteTeamMemberRequest) (*model.TeamMember, error) {
	t, err := GetTeamMember(ctx, &model.GetTeamMemberRequest{
		TeamID:       req.TeamID,
		TeamMemberID: req.TeamMemberID,
	})
	if err != nil {
		return nil, err
	}

	if err := db.DB.WithContext(ctx).Delete(t).Error; err != nil {
		log.Errorf("Failed to delete team member: %s", err)
		return nil, errors.Internal()
	}

	return t, nil
}
