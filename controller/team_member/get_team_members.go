package team_member

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/paginator"
	"github.com/ukasyah-dev/common/validator"
)

func GetTeamMembers(ctx context.Context, req *model.GetTeamMembersRequest) (*model.GetTeamMembersResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	tx := db.DB.WithContext(ctx).Model(&model.TeamMember{})

	data, pagination, err := paginator.Paginate[model.TeamMember](tx, &req.PaginationRequest)
	if err != nil {
		return nil, err
	}

	return &model.GetTeamMembersResponse{
		PaginationResponse: *pagination,
		Data:               data,
	}, nil
}
