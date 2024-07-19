package rpc

import (
	"context"

	"github.com/ukasyah-dev/authority-service/controller/permission"
	"github.com/ukasyah-dev/authority-service/model"
	pb "github.com/ukasyah-dev/pb/authority"
)

func (s *Server) CheckPermission(ctx context.Context, req *pb.CheckPermissionRequest) (*pb.CheckPermissionResponse, error) {
	res, err := permission.CheckPermission(ctx, &model.CheckPermissionRequest{
		ActionID: req.ActionID,
		TeamID:   req.TeamID,
		UserID:   req.UserID,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CheckPermissionResponse{
		Allowed: res.Allowed,
	}, nil
}
