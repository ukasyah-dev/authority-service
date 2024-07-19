package permission

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/paginator"
	"github.com/ukasyah-dev/common/validator"
)

func GetPermissions(ctx context.Context, req *model.GetPermissionsRequest) (*model.GetPermissionsResponse, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	req.PaginationRequest.Keys = []string{"ID"}

	tx := db.DB.WithContext(ctx).Model(&model.Permission{})

	data, pagination, err := paginator.Paginate[model.Permission](tx, &req.PaginationRequest)
	if err != nil {
		return nil, err
	}

	return &model.GetPermissionsResponse{
		PaginationResponse: *pagination,
		Data:               data,
	}, nil
}
