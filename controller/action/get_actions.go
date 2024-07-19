package action

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/paginator"
	"github.com/ukasyah-dev/common/validator"
)

func GetActions(ctx context.Context, req *model.GetActionsRequest) (*model.GetActionsResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	req.PaginationRequest.Keys = []string{"Name"}

	tx := db.DB.WithContext(ctx).Model(&model.Action{})

	data, pagination, err := paginator.Paginate[model.Action](tx, &req.PaginationRequest)
	if err != nil {
		return nil, err
	}

	return &model.GetActionsResponse{
		PaginationResponse: *pagination,
		Data:               data,
	}, nil
}
