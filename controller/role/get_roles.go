package role

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/paginator"
	"github.com/ukasyah-dev/common/validator"
)

func GetRoles(ctx context.Context, req *model.GetRolesRequest) (*model.GetRolesResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	tx := db.DB.WithContext(ctx).Model(&model.Role{})

	data, pagination, err := paginator.Paginate[model.Role](tx, &req.PaginationRequest)
	if err != nil {
		return nil, err
	}

	return &model.GetRolesResponse{
		PaginationResponse: *pagination,
		Data:               data,
	}, nil
}
