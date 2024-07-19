package team

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/constant"
	"github.com/ukasyah-dev/common/paginator"
	"github.com/ukasyah-dev/common/validator"
)

func GetTeams(ctx context.Context, req *model.GetTeamsRequest) (*model.GetTeamsResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	tx := db.DB.WithContext(ctx).Model(&model.Team{})

	userID, _ := ctx.Value(constant.UserID).(string)

	if userID != "" {
		tx = tx.Joins("INNER JOIN team_members tm ON tm.team_id = teams.id AND tm.user_id = ?", userID)
	}

	data, pagination, err := paginator.Paginate[model.Team](tx, &req.PaginationRequest)
	if err != nil {
		return nil, err
	}

	return &model.GetTeamsResponse{
		PaginationResponse: *pagination,
		Data:               data,
	}, nil
}
