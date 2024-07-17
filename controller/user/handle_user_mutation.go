package user

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/common/constant"
	"github.com/ukasyah-dev/common/log"
	commonModel "github.com/ukasyah-dev/common/model"
	identityModel "github.com/ukasyah-dev/identity-service/model"
)

func HandleUserMutation(ctx context.Context, mutation *commonModel.Mutation[identityModel.User]) (*commonModel.Empty, error) {
	if mutation.Type == constant.MutationCreated {
		return CreateUser(ctx, mutation.Data)
	} else if mutation.Type == constant.MutationUpdated {
		return UpdateUser(ctx, mutation.Data)
	} else if mutation.Type == constant.MutationDeleted {
		return DeleteUser(ctx, mutation.Data)
	}

	return &commonModel.Empty{}, nil
}

func CreateUser(ctx context.Context, req *identityModel.User) (*commonModel.Empty, error) {
	if err := db.DB.WithContext(ctx).Create(req).Error; err != nil {
		log.Errorf("Failed to create user: %s", err)
		return nil, err
	}

	return &commonModel.Empty{}, nil
}

func UpdateUser(ctx context.Context, req *identityModel.User) (*commonModel.Empty, error) {
	if err := db.DB.WithContext(ctx).Save(req).Error; err != nil {
		log.Errorf("Failed to update user: %s", err)
		return nil, err
	}

	return &commonModel.Empty{}, nil
}

func DeleteUser(ctx context.Context, req *identityModel.User) (*commonModel.Empty, error) {
	if err := db.DB.WithContext(ctx).Delete(req).Error; err != nil {
		log.Errorf("Failed to delete user: %s", err)
		return nil, err
	}

	return &commonModel.Empty{}, nil
}
