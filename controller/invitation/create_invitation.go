package invitation

import (
	"context"

	"github.com/ukasyah-dev/authority-service/db"
	"github.com/ukasyah-dev/authority-service/model"
	"github.com/ukasyah-dev/common/errors"
	"github.com/ukasyah-dev/common/id"
	"github.com/ukasyah-dev/common/log"
	"github.com/ukasyah-dev/common/validator"
)

func CreateInvitation(ctx context.Context, req *model.CreateInvitationRequest) (*model.Invitation, error) {
	if err := validator.Validate(req); err != nil {
		return nil, err
	}

	inv := &model.Invitation{
		ID:     id.New(),
		Email:  req.Email,
		RoleID: req.RoleID,
		TeamID: req.TeamID,
		Token:  id.New(56),
	}

	if err := db.DB.WithContext(ctx).Create(inv).Error; err != nil {
		log.Errorf("Failed to create invitation: %s", err)
		return nil, errors.Internal()
	}

	return inv, nil
}
