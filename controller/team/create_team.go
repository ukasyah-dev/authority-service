package team

import (
	"context"

	"github.com/ukasyah-dev/authority-service/controller/team_member"
	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/constant"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/id"
	"github.com/ukasyah-dev/common/log"
	"github.com/ukasyah-dev/common/validator"
)

func CreateTeam(ctx context.Context, req *model.CreateTeamRequest) (*model.Team, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	team := &model.Team{
		ID:   id.New(),
		Name: req.Name,
	}

	if err := db.DB.Create(team).Error; err != nil {
		log.Errorf("Failed to create team: %s", err)
		return nil, errors.Internal()
	}

	userID, _ := ctx.Value(constant.UserID).(string)

	if userID != "" {
		_, err := team_member.CreateTeamMember(ctx, &model.CreateTeamMemberRequest{
			RoleID: "admin",
			TeamID: team.ID,
			UserID: userID,
		})
		if err != nil {
			return nil, err
		}
	}

	return team, nil
}
