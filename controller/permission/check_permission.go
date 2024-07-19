package permission

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/log"
)

func CheckPermission(ctx context.Context, req *model.CheckPermissionRequest) (*model.CheckPermissionResponse, error) {
	var matched int64

	log.Debugf("Checking permission: actionId=%s teamId=%s userId=%s", req.ActionID, req.TeamID, req.UserID)

	sql := `SELECT COUNT(p.*) FROM permissions p
	INNER JOIN team_members tm ON tm.role_id = p.role_id
	WHERE p.action_id = ? AND tm.team_id = ? AND tm.user_id = ?`
	err := db.DB.WithContext(ctx).
		Raw(sql, req.ActionID, req.TeamID, req.UserID).
		Scan(&matched).Error
	if err != nil {
		log.Errorf("Failed to check permission: %s", err)
		return nil, errors.Internal()
	}

	return &model.CheckPermissionResponse{
		Allowed: matched > 0,
	}, nil
}
