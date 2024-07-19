package rpc

import (
	"context"

	"github.com/ukasyah-dev/authority-service/rpc/schema"
)

type Server struct {
	schema.UnimplementedAuthorityServer
}

func (s *Server) CheckPermission(ctx context.Context, req *schema.CheckPermissionRequest) (*schema.CheckPermissionResponse, error) {
	return &schema.CheckPermissionResponse{
		Allowed: false,
	}, nil
}
