package permission

import (
	"context"

	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/errors"
)

func CheckPermission(ctx context.Context, req *model.CheckPermissionRequest) (*model.CheckPermissionResponse, error) {
	return nil, errors.Internal("Not implemented")
}
